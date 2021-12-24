package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	////////////////
	go Talk("hello")
	go Talk("hi")
	////////////////
	fmt.Println("Waiting")
	time.Sleep(5 * time.Second)
	fmt.Println("END")
}

func Talk(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Println(prefix, i)
		time.Sleep(time.Second)
	}
}
