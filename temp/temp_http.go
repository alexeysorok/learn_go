package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string) {
	// response, err := http.Get("https://example.com")
	fmt.Print("Getting ", url, " ")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()                // освобождение сетевого подключения
	body, err := ioutil.ReadAll(response.Body) // читает все данные ответа
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}

func main() {

	responseSize("https://example.com") // use gorutine
	responseSize("https://ya.ru")
	responseSize("https://yandex.ru")
	responseSize("https://golang.org")
	responseSize("https://golang.org/doc")
	// time.Sleep(5 * time.Second)
}
