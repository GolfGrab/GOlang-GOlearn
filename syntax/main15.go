package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("myfile.txt")
	if err != nil {
		log.Println(err)
		return
	}
	//////////////////
	defer f.Close()
	//////////////////

	f.WriteString("Hello")
}
