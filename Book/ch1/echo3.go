// Использование менее ресуркоемкой Join
package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0:]) // вывести имя выполняемой команды
}
