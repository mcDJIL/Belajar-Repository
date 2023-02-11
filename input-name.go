package main

import "fmt"

func main() {
	var firstname string
	fmt.Println("Enter your firstname")
	fmt.Scanln(&firstname)

	var lastname string
	fmt.Println("Enter your lastname")
	fmt.Scanln(&lastname)

	fmt.Print("Namamu adalah " + firstname + " " + lastname)
}
