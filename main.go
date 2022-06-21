package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var text string
	fmt.Println("Введите название файла полностью")
	fmt.Scanf("%s", &text)

	doc, err := os.Open(text)
	if err != nil {
		panic("Файл не найден (попробуй указать вместе с расширением)")
	}
	defer doc.Close()

	res, err := os.Create("res.txt")
	if err != nil {
		panic("Невозможно создать файл (возможно проблемы с правами доступа)")
	}
	defer res.Close()

	data, err := ioutil.ReadAll(doc)
	if err != nil {
		panic("Невозможно считать содержимое файла (пиши мне)")
	}
	dataTxt := string(data)

	textCmp := "Балл: 0/1"
	textCmp2 := "Ответ (неправильный): \r\nН/Д"
	textCmp3 := "Ответ (неправильный): \r\n-1--1"

	ln := 0
	for i, el := range dataTxt {
		if len(dataTxt) < i+len(textCmp) || len(dataTxt) < i+len(textCmp2) || len(dataTxt) < i+len(textCmp3) {
			res.WriteString(string(el))
			continue
		}
		dt1 := dataTxt[i : i+len(textCmp)]
		dt2 := dataTxt[i : i+len(textCmp2)]
		dt3 := dataTxt[i : i+len(textCmp3)]
		if textCmp == dt1 {
			ln = len([]rune(textCmp)) + 1
		} else if textCmp2 == dt2 {
			ln = len([]rune(textCmp2)) + 1
		} else if textCmp3 == dt3 {
			ln = len([]rune(textCmp3)) + 1
		} else if ln != 0 {
			ln--
		} else {
			res.WriteString(string(el))
		}
	}
}
