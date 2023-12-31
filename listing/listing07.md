Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
```
и далее бесконечный вывод пустых значений типа int.
Объяснение:
1) Сначала программа создает два канала a и b, и генерирует в них значения из заданных наборов (1, 3, 5, 7 для a и 2, 4, 6, 8 для b).
2) Затем программа запускает функцию merge, которая объединяет в значения из каналов a и b в канал c, но не закрывает его. Запущенная в функции merge горутина с помощью конструкции select выбирает доступные значения из каналов a и b и отправляет их в канал c, но не закрывает этот канал после завершения работы.
3) После отправки все значений из каналов a и b в канал c программа продолжает выполнять цикл for v := range c. Но так как канал c не закрыт, то он цикл будет бесконечно ожидать новых значений в канале c.
4) После завершения работы горутин, которые генерировали значения в каналах a и b, и после того, как все значения были считаны из канала c, цикл for v := range c ожидает новых значений и выводит 0, так как канал c остается открытым и пустым.