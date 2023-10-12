package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestAppendLine(t *testing.T) {
	// Тестовые случаи с разными входными данными
	testCases := []struct {
		lines    []string
		line     string
		lineNums bool
		lineNum  int
		expected string
	}{
		{
			lines:    []string{},
			line:     "Тестовая строка",
			lineNums: false,
			lineNum:  0,
			expected: "Тестовая строка",
		},
		{
			lines:    []string{},
			line:     "Другая строка",
			lineNums: true,
			lineNum:  42,
			expected: "42: Другая строка",
		},
	}

	for _, testCase := range testCases {
		t.Run("AppendLine", func(t *testing.T) {
			actual := appendLine(testCase.lines, testCase.line, testCase.lineNums, testCase.lineNum)
			if len(actual) != 1 {
				t.Errorf("Ожидалась одна строка в результате, получено %d", len(actual))
			}

			if actual[0] != testCase.expected {
				t.Errorf("Неверный результат. Ожидалось: %s, получено: %s", testCase.expected, actual[0])
			}
		})
	}
}

func TestMyGrep(t *testing.T) {
	// Подготовка временного каталога и файлов с тестовыми данными
	inFile := filepath.Join("testdata", "input.txt")
	expectedFile := filepath.Join("testdata", "expected.txt")

	// Создание файлов с тестовыми данными
	err := os.WriteFile(inFile, []byte("Это простой тестовый текст.\nCъешь еще этих мягких французских булок, да выпей чаю.\nЕще не поздно.\nШирокая электрификация южных губерний даст мощный толчок подъему сельского хозяйства."), 0644)
	require.NoError(t, err)
	err = os.WriteFile(expectedFile, []byte("Cъешь еще этих мягких французских булок, да выпей чаю.\nЕще не поздно.\n"), 0644)
	require.NoError(t, err)

	// Запуск утилиты с заданными флагами и входными данными
	cmd := exec.Command("go", "run", "my_grep.go", "-i", "ЕЩЕ", inFile)

	actual, err := cmd.CombinedOutput()
	t.Logf("Результат выполнения утилиты:\n%s", string(actual))
	require.NoError(t, err)

	// Сравнение фактического вывода с ожидаемым
	expected, err := os.ReadFile(expectedFile)
	require.NoError(t, err)

	assert.Equal(t, string(expected), string(actual))
}
