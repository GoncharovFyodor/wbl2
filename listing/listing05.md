Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Объяснение:
В данной программе создается пользовательская структура customError, которая удовлетворяет интерфейсу error, и функция test(), которая возвращает указатель на nil типа *customError. Затем программа в функции main() объявляет переменную err типа error, присваивает ей результат вызова функции test(), и проверяет err.

Так как функция test() возвращает указатель на nil, то переменная err будет содержать указатель на nil. Условие err != nil будет истинным, так как указатель на nil указывает на nil, но сам не является nil, и программа выведет err.