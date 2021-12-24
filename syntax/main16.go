package main

import (
	"fmt"
)

func main() {
	fmt.Println("start.....")
	// do_fail_work()
	do_safe_work()
	fmt.Println("hello")
}

func do_fail_work() {
	panic("fail")
}

func do_safe_work() {
	defer func() {
		fmt.Println("defer")
		r := recover()
		if r != nil {
			fmt.Println("WORK FAIL :", r)
		}
	}()
	do_fail_work()
	fmt.Println("work success")
}
