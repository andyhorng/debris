package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

func main() {

	// fmt.Println(merge([]int{4, 5, 6, 7}, []int{1, 2, 3}))

	start := time.Now()
	cpu := 1 // runtime.NumCPU()

	runtime.GOMAXPROCS(cpu)

	list := rand.Perm(1000000)

	partSize := len(list) / cpu
	lists := make([][]int, cpu)

	for i := range lists {
		lists[i] = list[partSize*i : partSize*(i+1)]
	}

	c := make(chan []int)
	for i := range lists {
		go func(i int) {
			sort.Ints(lists[i])
			c <- lists[i]
		}(i)
	}

	for {
		a := <-c
		if len(a) == len(list) {
			// fmt.Println(a)
			break
		}
		b := <-c
		go func() {
			c <- merge(a, b)
		}()
	}

	fmt.Println(time.Since(start))
}
