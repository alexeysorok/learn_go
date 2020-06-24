package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type myStruct struct {
	myField int
}

// * значение в указателе
// & ссылка на указатель
func applyDicsount(s *subscriber) {
	s.rate = 4.99 // присвоение полю структуры по указателю
}

func main() {
	var s subscriber
	applyDicsount(&s)
	fmt.Println(s.rate)

	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	fmt.Println(pointer.myField) // обращение к полю по указателю
	pointer.myField = 10         // присвоение значения по указателю полю
	fmt.Println(pointer.myField)
}
