package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestChangeDir(t *testing.T) {
	// Создаем временную директорию для тестов
	tempDir := t.TempDir()
	err := os.Mkdir(tempDir+"/testdir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		args   []string
		errMsg string
	}{
		{[]string{"cd", tempDir}, ""},
		{[]string{"cd", "nonexistent"}, "chdir nonexistent: no such file or directory"},
		{[]string{"cd"}, "Использование: cd <директория>"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			err := changeDir(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}
		})
	}
}

func TestPrintWorkingDir(t *testing.T) {
	dir, err := printWorkingDir()
	if err != nil {
		t.Fatalf("Ошибка при получении текущей директории: %v", err)
	}
	t.Logf("Текущая директория: %s", dir)
}

func TestEcho(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
		errMsg   string
	}{
		{[]string{"echo", "Привет, мир!"}, "Привет, мир!", ""},
		{[]string{"echo"}, "", "Использование: echo <текст>"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			actual, err := echo(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}
			if actual != test.expected {
				t.Errorf("Ожидалось: %v, получено: %v", test.expected, actual)
			}
		})
	}
}

func TestKill(t *testing.T) {
	// В pid указывается имя процесса, который должен быть убит.
	pid := "12345"

	tests := []struct {
		args   []string
		errMsg string
	}{
		{[]string{"kill", pid}, ""},
		{[]string{"kill"}, "Использование: kill <PID>"},
		{[]string{"kill", "invalid"}, "Ошибка: неверный формат PID"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			err := kill(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}
		})
	}
}

func TestProcessList(t *testing.T) {
	out, err := processList()
	if err != nil {
		t.Fatalf("Ошибка при выполнении команды 'ps': %v", err)
	}
	t.Logf("Вывод 'ps':\n%s", out)
}

func TestExecuteExternalCommand(t *testing.T) {
	// Сохранение текущей директории
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Ошибка при получении текущей директории: %v", err)
	}

	// Установка временной директории для теста
	tempDir := t.TempDir()
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Ошибка при смене директории: %v", err)
	}

	defer func() {
		// Возвращаемся в исходную директорию после теста
		err = os.Chdir(originalDir)
		if err != nil {
			t.Fatalf("Ошибка при возвращении в исходную директорию: %v", err)
		}
	}()

	tests := []struct {
		args     []string
		expected string
		errMsg   string
	}{
		{[]string{"echo", "Hello, World!"}, "Hello, World!\n", ""},
		{[]string{"ls", "-l"}, "total 0\n", ""},
		{[]string{"nonexistent-command"}, "", "exec: \"nonexistent-command\": executable file not found"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			actual, err := executeExternalCommand(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}

			if string(actual) != test.expected {
				t.Errorf("Ожидалось: %v, получено: %v", test.expected, string(actual))
			}
		})
	}
}
