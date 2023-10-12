package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpack(s string) string {
	var res strings.Builder
	escape := false

	for i := 0; i < len(s); i++ {
		if escape {
			// Если включен режим escape, то идет добавление символа в результат и выключение этого режима.
			res.WriteByte(s[i])
			escape = false
		} else if s[i] == '\\' {
			// Если обнаружен символ '\', то режим escape выключается
			escape = true
		} else if unicode.IsDigit(rune(s[i])) {
			// Если обнаружена цифра, то собираем все цифры, чтобы получить число повторений
			count := ""
			for j := i; j < len(s); j++ {
				if unicode.IsDigit(rune(s[j])) {
					count += string(s[j])
				} else {
					break
				}
			}
			num, _ := strconv.Atoi(count)
			if num > 0 {
				// Если число больше 0, то предыдущий символ добавляется к результату
				// нужное число раз (за исключением самого символа)
				if i > 0 {
					res.WriteString(strings.Repeat(string(s[i-1]), num-1))
				}
			}
			i += len(count) - 1 //Перемещение указателя i на последнюю цифру числа
		} else {
			//В противном случае происходит добавление текущего символа в результат
			res.WriteByte(s[i])
		}
	}
	return res.String()
}

func main() {
	// Тестовые строки
	testCases := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
	for _, testCase := range testCases {
		fmt.Printf("%s => %s\n", testCase, unpack(testCase))
	}
}
