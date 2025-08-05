package main

import "fmt"

var name string = "pavan"

func main() {

	fmt.Println("Hello world!! ", name)
	println("Hello EArth!!")

	//variables types
	// 3 ways
	// var , const , type

	var a int //declared but not initialzed
	a = 10
	fmt.Println("value of a", a)
	a = 20
	fmt.Println("value of a after change", a)

	const b = 30 //constant, cannot be changed
	fmt.Println("const value of b", b)
	// b = 45
	// fmt.Println("value of b after change", b) // this will cause an error

	c := 55.5
	fmt.Print("1st line \n")
	fmt.Printf("C value is %f type of c %T \n", c, c)
	fmt.Println("C value is ", c, "type of c %T ", c)

	type person struct {
		name string
		age  int
	}

	/*
		var p1 person
		p1.name = "savitha"
		p1.age = 26

		fmt.Printf("My name is %s . I am %d old", p1.name, p1.age)

		var p, q, r int = 10, 20, 30
		fmt.Printf("p %d q %d r %d", p, q, r)
	*/

	fmt.Println("after comment")
	//formats
	/*
		%v -> value
		%d
		%f
		%t
	*/
	fmt.Printf("value of name %v", name) //value of name name pavan,pavan
	fmt.Printf("value of name %#v \n", name)
	fmt.Printf("value of name %T \n", name) //value of name string

	pi := 3.142323

	fmt.Printf("pi value %.20f \n", pi)

	//%10
	fmt.Printf("percentage is %%%d\n", 10)

	//task print binary of 12
	fmt.Printf("Hex value is %0x\n",12)
	z :="Savitha"
	fmt.Printf("%-10s%s",z,z)

}
