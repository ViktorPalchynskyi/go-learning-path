package main

import (
	"context"
	"fmt"
)

func main()  {
	fmt.Println("Lesson 15")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := generate(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squares := square(ctx, nums)
	evens := filter(ctx, func(i int) bool {return i%2 == 0}, squares)
	for n := range evens {
		fmt.Println(n) // 4, 16, 36, 64, 100
	}
}

func generate(ctx context.Context, numbs ...int) <-chan int{
	out := make(chan int)
	go func ()  {
		defer close(out)
		for _, n := range numbs {
			select{
			case <-ctx.Done():
				return 
			default:
				out <- n
			}
		}	
	}()
	return out
}

func square(ctx context.Context, in <-chan int) <-chan int{
	out := make(chan int)
	go func ()  {
		defer close(out)
		for n := range in {
			select{
			case <-ctx.Done():
				return 
			default:
				out <- n * n
			}
		}	
	}()
	return out
}

func filter(ctx context.Context, predicate func(int) bool, in <-chan int) <-chan int{
	out := make(chan int)
	go func ()  {
		defer close(out)
		for n := range in {
			select{
			case <-ctx.Done():
				return 
			default:
				if predicate(n) {
					out <- n
				}
			}
		}	
	}()
	return out
}