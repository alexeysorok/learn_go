//Оператор defer откладывает выполнение функции до того момента, как произойдет возврат из окружающей функции.

//Аргументы отложенных вызовов вычисляются сразу же, но вызов функции не происходит до того, как произойдет возврат из окружающей функции.

package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println("world")
	//defer fmt.Println("world2")

	//fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
