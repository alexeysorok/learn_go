package main

import "fmt"

func recurse() {
	fmt.Println("Oh no I'm Stuck")
	recurse()
}

func main() {
	recurse()
}
