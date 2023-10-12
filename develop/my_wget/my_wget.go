package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Проверка передачи URL в аргументе командной строки
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run my_wget.go <URL>")
		return
	}

	// Получение URL из аргументов командной строки
	url := os.Args[1]

	// Вызов функции для загрузки сайта
	err := downloadSite(url)
	if err != nil {
		fmt.Println("Ошибка при загрузке сайта:", err)
		return
	}
}

// Загрузка сайта
func downloadSite(url string) error {
	// Открытие HTTP-запроса к указанному URL
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Создание файла index.html для сохранения сайта
	out, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer out.Close()

	// Копирование содержимого ответа на запрос в созданный файл
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}
	fmt.Println("Загружен сайт:", url)

	return nil
}
