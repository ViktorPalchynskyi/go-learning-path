package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
}

func benchmarkPool(ctx context.Context, connString string, maxConns int32, numRequests int) time.Duration {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal(err)
	}
	config.MaxConns = maxConns

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	fmt.Printf("\n[Pool maxConns=%-2d] Launching %d parallel queries (pg_sleep(0.1))...\n", maxConns, numRequests)

	var wg sync.WaitGroup
	results := make(chan string, numRequests)

	start := time.Now()

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			qStart := time.Now()
			var x int
			err := pool.QueryRow(ctx, "SELECT pg_sleep(0.1), 1").Scan(nil, &x)
			elapsed := time.Since(qStart)
			if err != nil {
				results <- fmt.Sprintf("  query %2d: ERROR %v", i+1, err)
			} else {
				results <- fmt.Sprintf("  query %2d: done in %s", i+1, elapsed.Round(time.Millisecond))
			}
		}()
	}

	wg.Wait()
	close(results)

	total := time.Since(start)

	for r := range results {
		fmt.Println(r)
	}

	return total
}

func main() {
	var cfg Config
	envconfig.MustProcess("", &cfg)

	ctx := context.Background()
	numRequests := 10

	fmt.Println("=== Benchmark: pool size vs execution time ===")
	fmt.Printf("Requests: %d, each sleeps 100ms\n", numRequests)
	fmt.Printf("Expected: maxConns=2  → %d requests / 2 = %d waves × 100ms ≈ %dms\n",
		numRequests, numRequests/2, (numRequests/2)*100)
	fmt.Printf("Expected: maxConns=10 → %d requests / 10 = 1 wave × 100ms ≈ 100ms\n", numRequests)

	d2 := benchmarkPool(ctx, cfg.DatabaseURL, 2, numRequests)
	d10 := benchmarkPool(ctx, cfg.DatabaseURL, 10, numRequests)

	fmt.Println()
	fmt.Println("=== Results ===")
	fmt.Printf("  maxConns=2  → %s\n", d2.Round(time.Millisecond))
	fmt.Printf("  maxConns=10 → %s\n", d10.Round(time.Millisecond))
	fmt.Printf("  Speedup:       %.1fx faster with maxConns=10\n", float64(d2)/float64(d10))
	fmt.Println()
	fmt.Println("Conclusion: with maxConns=2 queries queue up waiting for a free connection.")
	fmt.Println("            10 goroutines compete for 2 slots — execution proceeds in waves.")
	fmt.Println("            With maxConns=10 all queries start simultaneously.")

	os.Exit(0)
}
