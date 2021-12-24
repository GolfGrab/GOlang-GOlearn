package main

import "fmt"

type PERSON struct {
	name string
}

func main() {
	m := make((map[interface{}]interface{}))
	m[1] = "hello"
	m["name"] = "Golf"
	fmt.Println(m["x"])

	if x, ok := m["abc"]; ok {
		fmt.Println(x)
	}
}
