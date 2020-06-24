package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b", "c"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "e")
	fmt.Println(slice, len(slice))
	slice = append(slice, "e", "x")
	fmt.Println(slice, len(slice))

	floafSlice := make([]float64, 10)
	fmt.Println(floafSlice)
}
