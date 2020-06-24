package main

import "fmt"

func rec() {
	defer recover()
	fmt.Println("recover")
	panic("My panic!")
	defer recover()
}
func main() {

	rec()
}
