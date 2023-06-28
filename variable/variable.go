package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	// var f
	// error : missing variable type or initializationsyntax
	// fmt.Println(f)

	f := "short"
	fmt.Println(f)
}
