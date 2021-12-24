package main

import "fmt"

func main() {
	fmt.Print("input score: ")
	var score int
	fmt.Scanln(&score)

	if 100 >= score && score >= 50 {
		fmt.Println("pass")
	} else if score < 50 && score >= 0 {
		fmt.Println("fail")
	} else {
		fmt.Println("🙄")
	}
}
