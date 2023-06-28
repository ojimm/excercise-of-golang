package main

import "fmt"

func main() {

	// basic loop
	fmt.Println("loop i")
	i := 1
	for i <= 2 {
		fmt.Println(i)
		i++
	}

	// typecially for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// infinite loop
	for {
		fmt.Println("infinite loop")
		break
	}

	//use continuation
	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
