package main

import (
	"fmt"
	"log"
)

type State interface {
	addItem(int) error
	requestItem() error
	pay(money int) error
	dispenseItem() error
}

// Торговый автомат имеет 4 состояния: есть товар, товар запрошен, товар оплачен, нет товара
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemsCount int
	itemPrice  int
}

// Создание нового торгового автомата
func NewVendingMachine(itemsCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemsCount: itemsCount,
		itemPrice:  itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

// Запрос товара
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

// Добавление товара
func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

// Оплата
func (v *VendingMachine) pay(money int) error {
	return v.currentState.pay(money)
}

// Выдача товара
func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

// Установка состояния
func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

// Увеличение счетчика товар на заданную величину
func (v *VendingMachine) incrementItemsCount(count int) {
	fmt.Printf("Добавлено товаров: %d", count)
	v.itemsCount = v.itemsCount + count
}

// Состояние "Нет товара"
type noItemState struct {
	vendingMachine *VendingMachine
}

// Запрос товара
func (i *noItemState) requestItem() error {
	return fmt.Errorf("Товара нет в наличии")
}

// Добавление товара
func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemsCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) pay(money int) error {
	return fmt.Errorf("Товара нет в наличии")
}
func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("Товара нет в наличии")
}

// Состояние "Есть товар"
type hasItemState struct {
	vendingMachine *VendingMachine
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemsCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("Товара нет в наличии")
	}
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	fmt.Printf("Товар запрошен\n")
	return nil
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("Добавлено товаров: %d\n", count)
	i.vendingMachine.incrementItemsCount(count)
	return nil
}

func (i *hasItemState) pay(money int) error {
	return fmt.Errorf("Пожалуйста, выберите сначала товар")
}
func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("Пожалуйста, выберите сначала товар")
}

// Состояние "Товар запрошен"
type itemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("Товар уже запрошен")
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("Идет выдача товара")
}

func (i *itemRequestedState) pay(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Внесено недостаточно денег. Внесите %d", i.vendingMachine.itemPrice)
	}
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	fmt.Println("Оплата произведена")
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("Сначала оплатите товар")
}

// Состояние "Товар оплачен"
type hasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("Идет выдача товара")
}

func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Идет выдача товара")
}

func (i *hasMoneyState) pay(money int) error {
	return fmt.Errorf("Товара нет в наличии")
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Выдача товара")
	i.vendingMachine.itemsCount = i.vendingMachine.itemsCount - 1
	if i.vendingMachine.itemsCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

func main() {
	vendingMachine := NewVendingMachine(1, 10)
	err := vendingMachine.requestItem()
	logErrorIfExists(err)

	err = vendingMachine.pay(5)
	logErrorIfExists(err)

	err = vendingMachine.dispenseItem()
	logErrorIfExists(err)

	fmt.Println()
	err = vendingMachine.addItem(2)
	logErrorIfExists(err)

	fmt.Println()

	err = vendingMachine.requestItem()
	logErrorIfExists(err)

	err = vendingMachine.pay(10)
	logErrorIfExists(err)

	err = vendingMachine.dispenseItem()
	logErrorIfExists(err)
}

func logErrorIfExists(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
