package main

import "fmt"

func main() {
	type noHP int
	type single bool

	var contact noHP = 6282261984331
	var status single = false
	fmt.Println(contact, status)
}
