package main

import "fmt"

func main() {
	var counters [3]int
	counters[0]++
	fmt.Println(counters[0])
	counters[0]++
	fmt.Println(counters[0])

	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	var primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes[0], primes[2], primes[4])

	notesMikro := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primesMikro := [5]int{2, 3, 5, 7, 11}
	fmt.Println(notesMikro[1], primesMikro[2], notesMikro)

	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)

	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
	fmt.Println(len(notes)) // lenght arrays

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	for index, notes := range notes {
		fmt.Println(index, notes)
	}

	for _, notes := range notes {
		fmt.Println(notes)
	}

}
