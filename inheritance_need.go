package main

import (
	"fmt"
)

type MoveAction struct {
	speed int
}

func (m MoveAction) Do(start int) int {
	return start + m.speed
}

type Dog struct {
	pos int
	MoveAction
}

type Cat struct {
	pos int
	MoveAction
}

func main() {
	dog := Dog{pos: 0, MoveAction: MoveAction{5}}
	cat := Cat{pos: 0, MoveAction: MoveAction{3}}

	fmt.Println("Dog is moving")
	dog.pos = dog.MoveAction.Do(dog.pos)
	fmt.Println("Dog complete moving", dog.pos)

	fmt.Println("Cat is moving")
	cat.pos = cat.MoveAction.Do(cat.pos)
	fmt.Println("Cat complete moving", cat.pos)
}
