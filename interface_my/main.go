package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

// в функцию можем передавать интерфейс, по сути тот тип которые реализует интерфейс, не нужно
// под каждый тип писать свою функцию
func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toy Canary")
	toy.MakeSound()
	toy = Horn("Toy Blater")
	toy.MakeSound()

	play(Whistle("Toy Canary"))
	play(Horn("Toy Blaster"))
}
