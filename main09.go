package main

import (
	"fmt"
)

func main() {
	a := [8]int{}
	mutateArray(a)
	fmt.Println("f1 arr a = : ", a)

	mutateSlideArray(a[:])
	fmt.Println("f2 arr a = : ", a)

}

func mutateArray(a [8]int) {
	a[0] = 10
}

func mutateSlideArray(a []int) {
	a[0] = 20
}
