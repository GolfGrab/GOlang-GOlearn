package main

import "fmt"

func main() {
	checktype("Gopher")
	checktype(10)
	checktype(true)

}

func checktype(v interface{}) {
	switch p := v.(type) {
	case string:
		fmt.Println("string", "hello", p)
	case int:
		fmt.Println("int", p)
	case bool:
		fmt.Println("Bool", "!p")
	}

	p, ok := v.(string)
	if ok {
		fmt.Println(p)
	} else {
		fmt.Print("sleepy")
	}

}
