package main

import "fmt"

type person struct {
	name     string
	nickname string
}

func (p *person) renick() {
	p.nickname = "sss"
	fmt.Println("inside mutate", p)

}

func mutateP(p *person) {
	p.name = "Hacker"
	fmt.Println("inside mutate", p)
}

func main() {
	p1 := person{
		name:     "GOLFFFFFF",
		nickname: "GG",
	}
	fmt.Println(p1)

	var p2 person
	p2.name = "kddkkkdddd"
	p2.nickname = "kdd"
	fmt.Println(p2)

	mutateP(&p2)
	fmt.Println("outside mutate", p2)

	p2.renick()
	fmt.Println("outside mutate", p2)

}
