package main

import "fmt"

func main() {
	fmt.Println("Hello man...!!!")

	//noraml for loop
	for i := 1; i < 10; i++ {
		fmt.Println("Hello ", i)
	}
	println("--------------------------")

	a := 10
	//while loop
	for a >= 5 {
		fmt.Println("hai ", a)
		a--
	}

	println("-------------------------")
	n := 5
	for {
		if n >= 10 {
			break
		}

		fmt.Println("number is ", n)

	}
}
