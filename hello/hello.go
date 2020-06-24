package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	fmt.Println("Ahrr")

	s1 := []int{1, 2, 3}
	fmt.Println(len(s1))
	println(s1)

	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	for idx := range s1 {
		fmt.Println(s1[idx])

	}

}
