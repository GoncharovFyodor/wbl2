package main

import (
	"fmt"
	"sync"
	"time"
)

// Создание канала, который будет закрыт после указанной задержки
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

// Объединение нескольких каналов в один
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если входных каналов нет, создается и возвращается закрытый канал
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		// Если входной канал только один, то просто возвращаем его
		return channels[0]
	}

	// Создание результирующего канала
	resChan := make(chan interface{})

	// Создание sync.Once для закрытия результирующего канала
	var closeOnce sync.Once

	// Горутина для отслеживания завершения каналов
	go func() {
		var wg sync.WaitGroup
		for _, ch := range channels {
			wg.Add(1)
			go func(ch <-chan interface{}) {
				defer wg.Done()
				for v := range ch {
					resChan <- v
				}

				//При завершении работы resChan закрывается только один раз
				closeOnce.Do(func() {
					close(resChan)
				})
			}(ch)
		}
		wg.Wait()
	}()
	return resChan
}

func main() {
	start := time.Now()
	doneChannels := []<-chan interface{}{
		sig(time.Hour),
		sig(2 * time.Minute),
		sig(5 * time.Second),
	}

	//Слияние каналов в один
	resChan := or(doneChannels...)
	<-resChan
	fmt.Printf("Завершено после %v\n", time.Since(start))
}
