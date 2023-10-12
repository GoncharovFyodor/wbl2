package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			actualOutput := unpack(testCase.input)
			if actualOutput != testCase.expectedOutput {
				t.Errorf("Для входных данных '%s' получено '%s', ожидалось '%s'.",
					testCase.input, actualOutput, testCase.expectedOutput)
			}
		})
	}
}
