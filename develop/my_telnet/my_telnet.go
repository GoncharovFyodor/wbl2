package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Парсинг аргументов командной строки
	host := flag.String("host", "", "Хост для подключения")
	port := flag.String("port", "23", "Порт для подключения")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут для подключения")

	flag.Parse()

	// Адрес для подключения
	address := fmt.Sprintf("%s:%s", *host, *port)

	//Таймаут для подключения
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Не удалось подключиться к серверу: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Подключено к", address)

	//Запуск горутины для чтения данных из сокета и их вывода в STDOUT
	go handleClient(conn)

	fmt.Println("Начинаем записывать ввод в сокет и получать вывод...")

	// Читаем ввод пользователя с консоли и записываем в сокет
	buf := make([]byte, 1024) // буфер для чтения клиентских данных
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			os.Exit(1)
		}

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Ошибка записи в сокет:", err)
			os.Exit(1)
		}
	}
}

func handleClient(conn net.Conn) {
	buf := make([]byte, 1024) // буфер для чтения клиентских данных
	for {
		n, err := conn.Read(buf) // читаем из сокета в buf
		if err != nil {
			fmt.Println("Соединение разорвано")
			os.Exit(0)
		}
		fmt.Println(string(buf[:n]))
	}
}
