package main

import (
	"fmt"
	"time"
)

var (
	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func main() {
	withGO()
	withoutGO()
}

func withoutGO() {
	defer timer()()
	fmt.Println("sum arr1 =", sum(arr1))
	fmt.Println("sum arr2 =", sum(arr2))
}

func withGO() {
	defer timer()()
	chRes1 := make(chan int)
	chRes2 := make(chan int)

	go func() {
		chRes1 <- sum(arr1)
	}()

	go func() {
		chRes2 <- sum(arr2)
	}()

	fmt.Println("sum arr1 =", <-chRes1)
	fmt.Println("sum arr2 =", <-chRes2)
}

func timer() func() {
	t := time.Now()
	return func() {
		diff := time.Now().Sub(t)
		fmt.Println(diff)
	}
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
		time.Sleep(200 * time.Millisecond)
	}
	return s
}
