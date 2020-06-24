package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, chanel chan Page) {
	// response, err := http.Get("https://example.com")
	// fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()                // освобождение сетевого подключения
	body, err := ioutil.ReadAll(response.Body) // читает все данные ответа
	if err != nil {
		log.Fatal(err)
	}
	chanel <- Page{URL: url, Size: len(body)}

}

func main() {

	pages := make(chan Page)
	urls := []string{"https://example.com", "https://ya.ru", "https://yandex.ru", "https://golang.org"}
	// use gorutine and chanel
	// go responseSize("https://ya.ru", sizes)       // канал передается при каждом вызове
	// go responseSize("https://yandex.ru", sizes)
	// go responseSize("https://golang.org", sizes)
	// go responseSize("https://golang.org/doc", sizes)
	// fmt.Println(<-sizes) // будет столько возвратов канала. сколько было вызовов
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages // получаем Pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}
