package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		// Приглашение для ввода команды
		fmt.Print("MyShell> ")

		// Считывание команд
		input := readInput()

		// Проверка наличия команды выхода
		if input == "exit" {
			break
		}

		// Разделение ввода на команды по пайпу
		commands := strings.Split(input, "|")

		var out []byte
		var err error
		for _, cmdStr := range commands {
			// Удаление лишних пробелов
			cmdStr = strings.TrimSpace(cmdStr)

			// Разделение команды на аргументы
			args := strings.Fields(cmdStr)

			//Проверка и выполнение команд
			switch args[0] {
			case "cd":
				err = changeDir(args)
				if err != nil {
					fmt.Println("Ошибка смены директории:", err)
				}
			case "pwd":
				dir, err := printWorkingDir()
				if err != nil {
					fmt.Println("Ошибка получения директории:", err)
				} else {
					fmt.Println("Текущая директория:", dir)
				}
			case "echo":
				result, err := echo(args)
				if err != nil {
					fmt.Println("Ошибка вывода аргумента:", err)
				} else {
					fmt.Println(result)
				}
			case "kill":
				err := kill(args)
				if err != nil {
					fmt.Println("Ошибка завершения процесса:", err)
				}
			case "ps":
				ps, err := processList()
				if err != nil {
					fmt.Println("Ошибка вывода информации по процессам:", err)
				} else {
					fmt.Println(ps)
				}
			default:
				out, err = executeExternalCommand(args)
			}
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}
		if len(out) > 0 {
			fmt.Println(string(out))
		}
	}
}

// Считывание ввода пользователя
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}

// Смена директории
func changeDir(args []string) error {
	var err error
	if len(args) != 2 {
		err = fmt.Errorf("Использование: cd <директория>")
	} else {
		err = os.Chdir(args[1])
	}
	return err
}

// Показать путь до текущего каталога
func printWorkingDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Вывод аргумента в STDOUT
func echo(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("Использование: echo <текст>")
	}
	return strings.Join(args[1:], " "), nil
}

func kill(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Использование: kill <PID>")
	}
	pID := strings.TrimSpace(args[1])

	cmd := exec.Command("kill", "-9", pID)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// Вывод общей информации по запущенным процессам
func processList() (string, error) {
	cmd := exec.Command("ps")
	ps, err := cmd.Output()
	return string(ps), err
}

// Выполнение внешней команды
func executeExternalCommand(args []string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	return out, err
}
