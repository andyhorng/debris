package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(10)
	list := rand.Perm(9999999)

	start := time.Now()
	mergeSort(list)
	duration := time.Since(start)

	fmt.Println("normal: ", duration)

	start = time.Now()
	goMergeSort(list)
	duration = time.Since(start)

	fmt.Println("parallel: ", duration)
}
