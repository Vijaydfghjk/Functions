package main

import "fmt"

func add(a int) func() int {

	return func() int {

		return a + 2

	}

}

func big(a *int) func() int {

	return func() int {

		return *a + 2
	}

}

func main() {

	add1 := add(1)

	fmt.Println(add1())

	var a = 2
	abig := big(&a)
	fmt.Println(abig())

	s1 := [...]int{1, 3, 4, 5, 6, 7, 8}

	for _, v := range &s1 {

		fmt.Println(v)
	}

}
