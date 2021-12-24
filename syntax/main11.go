package main

import "fmt"

type PERSON struct {
	name string
}

func (p PERSON) Talk() {
	n := 3
	for i := 0; i < n; i++ {
		fmt.Println(i, "hello my name is", p.name)
	}
}

type CAT struct{}

func (CAT) Talk() {
	n := 3
	for i := 0; i < n; i++ {
		fmt.Println(i, "🐱🐱🐱")
	}
}

type DOG struct{}

func (*DOG) Talk() {
	n := 3
	for i := 0; i < n; i++ {
		fmt.Println(i, "🐶🐶🐶")
	}
}

type TALK_ABLE interface {
	Talk()
}

func TALK_WITH(t TALK_ABLE) {
	t.Talk()
}

func main() {
	p1 := PERSON{
		name: "GolfGrab",
	}
	TALK_WITH(p1)
	TALK_WITH(CAT{})
	TALK_WITH(&DOG{})
}
