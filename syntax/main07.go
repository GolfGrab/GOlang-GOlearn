package main

import "fmt"

func main() {
	A := 10
	fmt.Println(A)

	ptrA := &A
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	*ptrA = 200
	fmt.Println(A)

}
