package main

import (
	"fmt"
	"reflect"
)

func main() {
	//arithmetics operator and assignments operator
	var a = 10
	var b = 10
	a += b
	fmt.Println(a)

	var i = 5
	i += 5
	fmt.Println(i)

	//Logic operator
	averagePAS := 85
	averageAbsensi := 80

	lulusSMK := averagePAS >= 80 && averageAbsensi >= 80
	fmt.Println(lulusSMK)

	if lulusSMK {
		fmt.Println("Selamat anda lulus SMK")
	} else {
		fmt.Println("Maaf anda tidak lulus")
	}
	fmt.Println(reflect.TypeOf(lulusSMK))
}
