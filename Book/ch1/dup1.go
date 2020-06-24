// Выводит текст каждой строки, которая появляется в
// стандартном вводе более одного раза, а также кол-во
// ее появлений

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++ // line := input.Text()
		// counts[line] = counts[line] + 1 -- полная форма записи
	}
	//  Примечание: игнорируем потенциальные
	// ошибки из input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
