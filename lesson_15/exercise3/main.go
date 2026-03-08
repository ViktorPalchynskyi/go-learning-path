package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

func formatJSONWithPool(data map[string]any) string {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	json.NewEncoder(buf).Encode(data)
	return buf.String()
}

func formatJSONWithoutPool(data map[string]any) string {
	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(data)
	return buf.String()
}

func main() {
	fmt.Println("Lesson 15 — run: go test -bench=. -benchmem ./exercise3/")
}
