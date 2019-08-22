// coding problem: print number 0 - 100
// if number is divisile by 3, print fizz
// if number is divisible by 5, print buzz
// if number is divisible by 3 and 5, print fizz buzz

package main

import "fmt"

func fizzbuzz() {
	countTo := 100

	for i := 0; i <= countTo; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
