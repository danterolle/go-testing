package main

import "fmt"

func main() {
	var i, x int
	for i = 0; i <= 100; i++ {
		switch x {
		case i % 15:
			{
				fmt.Println(i, "FizzBuzz")
			}
		case i % 3:
			{
				fmt.Println(i, "Fizz")
			}
		case i % 5:
			{
				fmt.Println(i, "Buzz")
			}
		default:
			fmt.Println(i, "No one of the previous use cases")
		}
	}
}
