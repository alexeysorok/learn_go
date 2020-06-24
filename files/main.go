package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover() // если нет ошибок выходим из функции
	if p == nil {
		return
	}
	err, ok := p.(error) // если есть ошибки преобразуем к типу ошибки
	if ok {
		fmt.Println(err) // получаем и выводим ошибку
	} else {
		panic(p) // если значение паники не является признаком ошибки
	}

}

func scanDirectory(path string) {
	fmt.Println(path)
	// files, err := ioutil.ReadDir("C:\\Users\\admin\\go\\")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// return err
		panic(err) // инициируется паника и возвращаются все рекурсивные вызовы возвращают управление
	}
	for _, file := range files {
		filePatch := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePatch)
		} else {
			fmt.Println(filePatch)
		}
	}

}

func main() {
	defer reportPanic() // откладыаем функцию, так при панике не завершится программа
	scanDirectory("C:\\Users\\admin\\go\\")

}
