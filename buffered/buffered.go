package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var point []int
	input := make(chan []int)
	buffer := make([]int, 10)

	go process(5*time.Second, input)

	for i := 1; i <= 60; i++ {

		point = buffer[0 : len(point)+1]
		point[len(point)-1] = i

		if len(point) >= 10 {
			input <- point
			buffer = make([]int, 10)
			point = nil
		}

		t := rand.Int63n(999) + 1
		fmt.Println("tick", i, "sleep", t)
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
}

func process(flushAt time.Duration, input <-chan []int) {

	ticker := time.NewTimer(flushAt)

	for {
		select {
		case buffer := <-input:
			fmt.Println("trigger", buffer)
			ticker.Reset(flushAt)
			fmt.Println("reset timer")
		case <-ticker.C:
			fmt.Println("timer")
		}
	}
}
