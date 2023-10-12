package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	// Получение точного времени с сервера NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
	}

	//Вывод текущего времени
	fmt.Println("Точное время:", ntpTime.Format(time.RFC3339))
}
