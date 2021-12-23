package main

import "fmt"

func main() {
	dict := make(map[string]string)
	dict["one"] = "1"
	dict["two"] = "2"
	x, ok := dict["x"]
	fmt.Println(ok)
	fmt.Println(x)
	fmt.Println(dict)

	dict["x"] = "x now good"
	x, ok2 := dict["x"]
	fmt.Println(ok2)
	fmt.Println(x)
	fmt.Println(dict)

	if x, ok := dict["x"]; ok {
		fmt.Println(x)
		fmt.Println("hear 1")
	} else {
		fmt.Println("not hear 1")
	}

	delete(dict, "x")
	x, ok3 := dict["x"]
	fmt.Println(ok3)
	fmt.Println(x)
	fmt.Println(dict)

	if x, ok := dict["x"]; ok {
		fmt.Println(x)
		fmt.Println("hear 2")
	} else {
		fmt.Println("not hear 2")
	}

	for k, v := range dict {
		fmt.Println(k, v)
	}

	dict2 := map[string]string{
		"apple":  "🍎",
		"banana": "🍌",
		"cherry": "🍒",
		"orange": "🍊",
	}

	fmt.Println(dict2)
}
