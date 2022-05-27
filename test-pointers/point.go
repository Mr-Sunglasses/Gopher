package main

import "fmt"

func main() {

	myval := 20

	var ptr = &myval

	*ptr = *ptr + 2

	fmt.Println("The value of myval is ", myval)
	fmt.Println("The addreess of ptr is ", ptr)
	fmt.Println("The value of ptr is", *ptr)
}
