// выводит индекс
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for indx, arg := range os.Args[1:] {
		s += sep + arg + string(indx)
		sep = " "

	}
	fmt.Println(s)
}
