package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, c chan float64) (float64, chan float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	return x, c
}

func main() {
	start := time.Now()

	/*
		~15 seconds
			for i := 1; i < 15; i++ {
				n := fib(float64(i))
				fmt.Printf("Fib(%v): %v\n", i, n)
			}
	*/

	ch := make(chan float64, 15)
	// ~0.00013 seconds
	go fib(15, ch)

	for i := range ch {
		fmt.Println(i)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
