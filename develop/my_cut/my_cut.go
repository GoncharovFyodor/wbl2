package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "Выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "Использовать другой разделитель")
	separated := flag.Bool("s", false, "Только строки с разделителем")
	flag.Parse()

	// Обработка входных данных
	var in io.Reader
	if flag.NArg() == 0 {
		in = os.Stdin
	} else {
		fileName := flag.Arg(0)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		in = file
	}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}
		parts := strings.Split(line, *delimiter)
		if len(parts) == 1 && parts[0] == "" {
			continue
		}

		indices := strings.Split(*fields, ",")

		// Выбор полей (колонок) и вывод
		var selected []string
		for _, index := range indices {
			i, err := strconv.Atoi(index)
			if err != nil || i < 1 || i > len(parts) {
				continue
			}
			selected = append(selected, parts[i-1])
		}
		if len(selected) > 0 {
			fmt.Println(strings.Join(selected, *delimiter))
		}
	}
}
