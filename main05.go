package main

import "fmt"

func main() {

	//one
	var a [5]int

	a[0] = 1
	a[1] = 10
	a[2] = 100

	for i, v := range a {
		fmt.Println(i, v)
	}

	//two
	b := [5]int{1, 2, 3, 4, 5}

	for i, v := range b {
		fmt.Println(i, v)
	}

	//three
	c := make([]int, 5)

	c[0] = 1
	c[2] = 5
	c[3] = 7

	c = append(c, 10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
