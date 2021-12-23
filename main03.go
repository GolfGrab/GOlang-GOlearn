package main03

import "fmt"

func main03() {
	fmt.Print("input a fruit name: ")
	var fruit string
	fmt.Scanln(&fruit)

	if fruit == "" {
		fmt.Println("nooooo 😥")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("🍎")
	case "banana":
		fmt.Println("🍌")
	case "orange":
		fmt.Println("🍊")
	case "grape":
		fmt.Println("🍇")
	case "watermelon":
		fmt.Println("🍉")
	case "strawberry":
		fmt.Println("🍓")
	default:
		fmt.Println("👾")
	}
}
