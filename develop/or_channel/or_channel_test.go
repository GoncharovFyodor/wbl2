package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	t.Run("No channels", func(t *testing.T) {
		orChan := or()
		select {
		case _, ok := <-orChan:
			if ok {
				t.Fatal("or() должна закрыть канал сразу же")
			}
		case <-time.After(1 * time.Second):
			t.Fatal("Неожиданный таймаут")
		}
	})
	t.Run("Single channel", func(t *testing.T) {
		doneChan := sig(2 * time.Second)
		select {
		case <-or(doneChan):
		// Функция or не должна закрывать doneChan на протяжении 2 секунд
		case <-time.After(3 * time.Second):
			t.Fatal("or(doneChan) должна закрыться не раньше, чем через 2 секунды")
		}
	})

	t.Run("Multiple channels", func(t *testing.T) {
		doneChan1 := sig(2 * time.Second)
		doneChan2 := sig(1 * time.Second)
		doneChan3 := sig(3 * time.Second)
		orChan := or(doneChan1, doneChan2, doneChan3)
		select {
		case <-orChan:
		// Закроется по крайней мере один канал
		case <-time.After(4 * time.Second):
			t.Fatal("or(doneChan1, doneChan2, doneChan3) должна закрыться не раньше, чем через 3 секунды")
		}
	})
}
