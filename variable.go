package main

import (
	"fmt"
	"reflect"
)

func main() {

	firstName := "Moch "
	lastName := "Djauharil"

	fmt.Println(firstName + lastName)
	var (
		name   = "Aril"
		age    = "16"
		hobby  = "Coding"
		haveGF = "true"
	)
	fmt.Println("Perkenalkan nama saya adalah " + name + ", " + "umur saya adalah " + age + " tahun. " +
		"Hobi saya adalah " + hobby + ". " + "Apakah aril punya GF? " + haveGF)

	addition := 10 + 5
	fmt.Println(addition)
	subtraction := 100 - 100
	fmt.Println(subtraction)
	multiplication := 10 * 10
	fmt.Println(multiplication)
	modulus := 10 % 3
	fmt.Println(modulus)

	capek := false
	bosen := true
	if capek {
		fmt.Println("Tidurlah")
	} else if bosen {
		fmt.Println("Ngegamelah")
	} else {
		fmt.Println("Belajarlah")
	}

	const person = "Aril"
	fmt.Println(person)

	//mengecek apakah variabel tersebut string, dengan conditional statement
	var isString interface{} = "This is string"
	if _, ok := isString.(string); ok {
		fmt.Println("Yang anda masukan adalah string")
	} else {
		fmt.Println("Yang anda masukan bukan string")
	}

	//mengecek variable dengan typeof
	a := 10
	b := ""
	c := true
	d := 2.5
	e := ".gitignore"
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
}
