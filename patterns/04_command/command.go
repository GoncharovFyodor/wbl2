package main

import "fmt"

// Кнопка, с которой связана команда
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Интерфейс команды
type Command interface {
	execute()
}

// Команда выключения
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Команда включения
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("Включение ТВ")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("Выключение ТВ")
}

func main() {
	tv := &TV{}
	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}
	onButton := &Button{
		command: onCommand,
	}
	onButton.press()
	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
