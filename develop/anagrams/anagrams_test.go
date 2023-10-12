package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		in       []string
		expected map[string][]string
	}{
		{
			in: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
				"кот":    {"кот", "ток"},
			},
		},
		{
			in: []string{"абв", "вба", "где", "дег"},
			expected: map[string][]string{
				"абв": {"абв", "вба"},
				"где": {"где", "дег"},
			},
		},
		{
			in:       []string{"яблоко", "банан"},
			expected: map[string][]string{},
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			result := findAnagrams(testCase.in)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Ожидалось: %v, получено: %v", testCase.expected, result)
			}
		})
	}
}
