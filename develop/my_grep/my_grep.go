package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Флаги для настройки фильтрации
	after := flag.Int("A", 0, "Печатать N строк после совпадения (по умолчанию: 0)")
	before := flag.Int("B", 0, "Печатать N строк до совпадения (по умолчанию: 0)")
	context := flag.Int("C", 0, "Печатать +-N строк вокруг совпадения (по умолчанию: 0)")
	count := flag.Bool("с", false, "Печатать количество строк")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	invert := flag.Bool("v", false, "Вместо совпадения исключать")
	fixed := flag.Bool("F", false, "Точное совпадение со строкой, не паттерн")
	lineNums := flag.Bool("n", false, "Напечатать номер строки")

	flag.Parse()

	// Проверка наличия аргумента - паттерна для поиска
	if flag.NArg() < 1 {
		fmt.Println("Использование: mygrep [опции] паттерн")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Считывание паттерна из аргумента командной строки
	pattern := flag.Arg(0)

	// Открытие файла для чтения или использование stdin
	var in io.Reader
	if flag.NArg() == 2 {
		fileName := flag.Arg(1)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Ошибка при открытии файла: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		in = file
	} else {
		in = os.Stdin
	}

	scanner := bufio.NewScanner(in)
	matchingLines := 0
	var lines []string
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		//Применение флага -F (точное совпадение)
		if *fixed {
			for _, line := range lines {
				if strings.Contains(line, pattern) != *invert {
					printLine(line, *lineNums, lineNum)
					matchingLines++
				}
			}
		} else {
			// Применение флага -i (игнорирование регистра)
			if *ignoreCase {
				lineLower := strings.ToLower(line)
				patternLower := strings.ToLower(pattern)
				if strings.Contains(lineLower, patternLower) != *invert {
					printLine(line, *lineNums, lineNum)
					matchingLines++
				}
			} else {
				// Поиск совпадения с учетом флага -v
				if (strings.Contains(line, pattern) && !*invert) || (!strings.Contains(line, pattern) && *invert) {
					printLine(line, *lineNums, lineNum)
					matchingLines++
				}
			}
		}

		// Печатать +N строк после совпадения (если выставлен соответствующий флаг)
		if *after > 0 && matchingLines > 0 && matchingLines <= *after {
			lines = appendLine(lines, line, *lineNums, lineNum)
		}

		// Печатать +N строк до совпадения (если выставлен соответствующий флаг)
		if *before > 0 && matchingLines > 0 && matchingLines <= *before+1 {
			lines = appendLine(lines, line, *lineNums, lineNum)
		}

		// Печатать +-N строк вокруг совпадения (если выставлен соответствующий флаг)
		if *context > 0 && matchingLines > 0 && matchingLines <= *context*2+1 {
			lines = appendLine(lines, line, *lineNums, lineNum)
		}

		// Удаление старых строк
		if len(lines) > *context*2+1 {
			lines = lines[1:]
		}

		// Печать результата
		for _, line := range lines {
			fmt.Println(line)
		}

		// Печать количества совпадений, если флаг -c установлен
		if *count {
			fmt.Println("Количество совпадений:", matchingLines)
		}

	}
}

// Добавление строки с учетом номера
func appendLine(lines []string, line string, lineNums bool, num int) []string {
	if lineNums {
		lines = append(lines, fmt.Sprintf("%d: %s", num, line))
	} else {
		lines = append(lines, line)
	}
	return lines
}

// Печать строки с учетом номера
func printLine(line string, lineNums bool, num int) {
	if lineNums {
		fmt.Printf("%d: %s\n", num, line)
	} else {
		fmt.Println(line)
	}
}
