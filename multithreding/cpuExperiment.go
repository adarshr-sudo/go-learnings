package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	for i:=0; i < 1_000_000_000; i++ {
	}

	fmt.Printf("worker %d finished in %v\n", id, time.Since(start))
}

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(2))

	var wg sync.WaitGroup

	wg.Add(4)

	start := time.Now()

	for i:=1;i <= 4; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Total: ", time.Since(start))
}
