package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestCutUtility(t *testing.T) {
	// Подготовка временного файла с тестовыми данными
	inData := "1\tИван\tПетров\n2\tСергей\tСидоров"
	expected := "1\tПетров\n2\tСидоров\n"
	inFile := "input.txt"

	err := os.WriteFile(inFile, []byte(inData), 0644)
	if err != nil {
		t.Fatalf("Ошибка создания файла с тестовыми данными: %v", err)
	}

	defer func() {
		err := os.Remove(inFile)
		if err != nil {
			t.Fatalf("Ошибка удаления временного файла: %v", err)
		}
	}()

	// Запуск утилиты с заданными флагами и входными данными
	cmd := exec.Command("go", "run", "my_cut.go", "-f", "1,3", inFile)
	actual, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Ошибка выполнения утилиты: %v", err)
	}

	// Сравнение фактического вывода с ожидаемым
	if string(actual) != expected {
		t.Errorf("Ожидаемый вывод: %s, Фактический вывод: %s", expected, string(actual))
	}
}
