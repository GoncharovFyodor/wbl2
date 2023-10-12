package main

import (
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	//Временная карта анаграмм
	tmp := make(map[string][]string)

	for _, word := range words {
		// Приведение слова к нижнему регистру и сортировка букв
		word = strings.ToLower(word)
		letters := strings.Split(word, "")
		sort.Strings(letters)
		sortedWord := strings.Join(letters, "")

		// Добавление слова в соответствующее множество анаграмм
		if _, found := tmp[sortedWord]; !found {
			tmp[sortedWord] = []string{word}
		} else {
			tmp[sortedWord] = append(tmp[sortedWord], word)
		}
	}

	unique := make(map[string]bool)

	//Удаление множеств из одного элемента
	for key, value := range tmp {
		if len(value) < 2 {
			delete(tmp, key)
		}
		for i := range value {
			if !unique[value[i]] {
				unique[value[i]] = true
			} else {
				value[i] = value[len(value)-1]
				tmp[key] = value[:len(value)-1]
			}
		}
	}

	anagrams := make(map[string][]string, len(tmp))

	for _, value := range tmp {
		sort.Strings(value)
		anagrams[value[0]] = value
	}

	return anagrams
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток"}
	anagrams := findAnagrams(words)
	for key, val := range anagrams {
		println(key, ":", strings.Join(val, ", "))
	}
}
