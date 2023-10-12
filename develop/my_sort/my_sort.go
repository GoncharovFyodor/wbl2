package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Line Структура для хранения строк и их ключенй
type Line struct {
	str string
	key string
}

func main() {
	// Флаги для настройки сортировки
	inFileName := flag.String("i", "", "Имя входного файла")
	outFileName := flag.String("o", "", "Имя выходного файла")
	keySortColumn := flag.Int("k", 0, "Колонка для сортировки (по умолчанию: 0)")
	numeric := flag.Bool("n", false, "Сортировка по числовому значению")
	rev := flag.Bool("r", false, "Сортировка в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	monthSort := flag.Bool("M", false, "Сортировка по названию месяца")
	ignoreTrailingSpace := flag.Bool("b", false, "Игнорировать хвостовые пробелы")
	checkSorted := flag.Bool("c", false, "Проверять, отсортированы ли данные")
	numericSuffix := flag.Bool("h", false, "Сортировка по числовому значению с учетом суффиксов")

	flag.Parse()

	if *inFileName == "" {
		fmt.Println("Не указано имя входного файла")
		return
	}

	// Открытие входного файла
	inFile, err := os.Open(*inFileName)
	if err != nil {
		fmt.Printf("Ошибка при открытии входного файла: %v\n", err)
		return
	}
	defer inFile.Close()

	lines := make([]Line, 0)

	// Чтение строк из файла и возвращение ключей
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, Line{str: text, key: extractKey(text, *keySortColumn, *monthSort, *numericSuffix)})
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении входного файла: %v\n", err)
		return
	}

	// Проверка отсортированности данных (если установлен соответствующий флаг)
	if *checkSorted && !isSorted(lines, *rev) {
		fmt.Println("Данные не отсортированы.")
		return
	}

	// Сортировка строк в соответствии с флагами
	if *numericSuffix {
		sort.SliceStable(lines, func(i, j int) bool {
			return compareNumbersWithSuffix(lines[i].key, lines[j].key, *rev)
		})
	} else {
		sort.SliceStable(lines, func(i, j int) bool {
			return compareStrings(lines[i].key, lines[j].key, *numeric, *rev, *ignoreTrailingSpace)
		})
	}

	// Удаление дубликатов, если установлен unique
	if *unique {
		lines = removeDuplicates(lines)
	}

	// Определение выходного файла, если указан флаг outFileName
	outFile := os.Stdout
	if *outFileName != "" {
		outFile, err = os.Create(*outFileName)
		if err != nil {
			fmt.Printf("Ошибка при создании выходного файла: %v\n", err)
			return
		}
		defer outFile.Close()
	}

	// Запись отсортированных строк в выходной файл или вывод на консоль
	for _, line := range lines {
		fmt.Fprintln(outFile, line.str)
	}
}

// Удаление дубликатов
func removeDuplicates(lines []Line) []Line {
	allKeys := make(map[Line]bool)
	list := []Line{}
	for _, line := range lines {
		if _, value := allKeys[line]; !value {
			allKeys[line] = true
			list = append(list, line)
		}
	}
	return list
}

// Сравнение строк с ученом различных флагов
func compareStrings(a, b string, numeric bool, rev bool, ignoreTrailingSpace bool) bool {
	if ignoreTrailingSpace {
		a = strings.Trim(a, " ")
		b = strings.Trim(b, " ")
	}
	if numeric {
		nA, errA := strconv.Atoi(a)
		nB, errB := strconv.Atoi(b)
		if errA != nil && errB != nil {
			if rev {
				return nA < nB
			}
			return nA > nB
		}
	}

	if rev {
		return a < b
	}
	return a > b
}

// Сравнение строк с числами и суффиксами
func compareNumbersWithSuffix(a string, b string, rev bool) bool {
	numA, sufA := extractSuffix(a)
	numB, sufB := extractSuffix(b)
	if numA == numB {
		if rev {
			return sufA > sufB
		}
		return sufA < sufB
	}
	if rev {
		return numA > numB
	}
	return numA < numB
}

// Извлечение суффикса из числа
func extractSuffix(s string) (int, string) {
	parts := strings.SplitN(s, ".", 2)
	num := 0
	if len(parts) > 0 {
		if n, err := strconv.Atoi(parts[0]); err == nil {
			num = n
		}
	}

	suffix := ""
	if len(parts) > 1 {
		suffix = parts[1]
	}
	return num, suffix
}

// Извлечение ключа из строки на основе заданных параметров
func extractKey(text string, keySortColumn int, monthSort bool, numericSuffix bool) string {
	fields := strings.Fields(text)
	if keySortColumn >= len(fields) {
		return ""
	}
	key := fields[keySortColumn]

	if monthSort {
		time, err := time.Parse("January", key)
		if err == nil {
			return time.Format("01")
		}
	}

	if numericSuffix {
		parts := strings.SplitN(key, ".", 2)
		if len(parts) == 2 {
			numPart, err := strconv.Atoi(parts[0])
			if err == nil {
				return fmt.Sprintf("%06d.%s", numPart, parts[1])
			}
		}
	}

	return key
}

// Проверка, отсортированы ли строки
func isSorted(lines []Line, rev bool) bool {
	for i := 1; i < len(lines); i++ {
		if (rev && lines[i].key > lines[i-1].key) || (!rev && lines[i].key < lines[i-1].key) {
			return false
		}
	}
	return true
}
