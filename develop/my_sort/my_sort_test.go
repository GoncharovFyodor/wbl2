package main

import (
	"testing"
)

func TestCompareNumbersWithSuffix(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		rev      bool
		expected bool
	}{
		{"123.456", "654.321", false, true},  // Сравнение чисел с суффиксами
		{"123.456", "654.321", true, false},  // Сравнение чисел с суффиксами
		{"123.456", "123.654", false, true},  // Сравнение чисел с суффиксами
		{"123.654", "123.456", false, false}, // Сравнение чисел с суффиксами
		{"123.xyz", "123.abc", false, false}, // Сравнение строк с суффиксами
	}
	for _, test := range tests {
		actual := compareNumbersWithSuffix(test.a, test.b, test.rev)
		if actual != test.expected {
			t.Errorf("Строка 1: %s, строка 2: %s. Ожидалось: %v, получено: %v", test.a, test.b, test.expected, actual)
		}
	}
}

func TestCompareStrings(t *testing.T) {
	tests := []struct {
		a                   string
		b                   string
		numeric             bool
		rev                 bool
		ignoreTrailingSpace bool
		expected            bool
	}{
		{"123", "456", true, false, false, true},          //Сравнение чисел
		{"123", "456", true, true, false, false},          //Сравнение чисел (в обратном порядке)
		{"яблоко", "банан", false, false, false, true},    //Сравнение строк
		{"яблоко", "банан", false, true, false, false},    //Сравнение строк (в обратном порядке)
		{"яблоко", "яблоко ", false, false, true, true},   //Сравнение строк с игнорированием пробелов
		{"яблоко ", "яблоко  ", false, false, true, true}, //Сравнение строк с игнорированием пробелов и разной длиной
	}
	for _, test := range tests {
		actual := compareStrings(test.a, test.b, test.numeric, test.rev, test.ignoreTrailingSpace)
		if actual != test.expected {
			t.Errorf("Строка 1: %s, строка 2: %s. Ожидалось: %v, получено: %v", test.a, test.b, test.expected, actual)
		}
	}
}

func TestExtractKey(t *testing.T) {
	tests := []struct {
		text          string
		keySortColumn int
		monthSort     bool
		numericSuffix bool
		expectedKey   string
	}{
		{"яблоко 123 January", 1, false, false, "123"},   // Тест с числловым ключом и месяцем
		{"яблоко 123.456", 1, false, true, "000123.456"}, // Тест с числовым ключом и суффиксом
		{"яблоко March", 1, true, false, "03"},           // Тест с месяцем
		{"яблоко", 1, false, false, ""},                  // Тест без ключа
	}
	for _, test := range tests {
		key := extractKey(test.text, test.keySortColumn, test.monthSort, test.numericSuffix)
		if key != test.expectedKey {
			t.Errorf("Строка: %s. Ожидался ключ: %s, получено: %s", test.text, test.expectedKey, key)
		}
	}
}

func TestIsSorted(t *testing.T) {
	lines := []Line{
		{str: "ананас", key: "ананас"},
		{str: "банан", key: "банан"},
		{str: "яблоко", key: "яблоко"},
	}
	tests := []struct {
		lines    []Line
		rev      bool
		expected bool
	}{
		{lines, false, true}, // Строки отстортированы по возрастанию
		{lines, true, false}, // Строки отстортированы по убыванию
	}
	for _, test := range tests {
		actual := isSorted(test.lines, test.rev)
		if actual != test.expected {
			t.Errorf("Строки: %v, обратный порядок: %v. Ожидалось: %v, получено: %v", test.lines, test.rev, test.expected, actual)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	lines := []Line{
		{str: "ананас", key: "ананас"},
		{str: "банан", key: "банан"},
		{str: "яблоко", key: "яблоко"},
		{str: "банан", key: "банан"},
		{str: "яблоко", key: "яблоко"},
	}
	expected := []Line{
		{str: "ананас", key: "ананас"},
		{str: "банан", key: "банан"},
		{str: "яблоко", key: "яблоко"},
	}

	actual := removeDuplicates(lines)

	if len(actual) != len(expected) {
		t.Errorf("Строки: %v. Ожидалось: %v, получено: %v", lines, expected, actual)
	}
}
