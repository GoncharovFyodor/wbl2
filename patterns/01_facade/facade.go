package main

import (
	"fmt"
	"log"
)

// Подсистема инвентаря
type InventorySystem struct{}

// Проверить инвентарь
func (i *InventorySystem) CheckInventory(productId string, quantity int) error {
	fmt.Printf("Проверка склада на наличие товара \"%s\", кол-во: %d\n", productId, quantity)
	return nil
}

// Подсистема платежного шлюза
type PaymentGateway struct{}

// Обработка платежа
func (p *PaymentGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Обработка платежа на сумму %f\n", amount)
	return nil
}

// Подсистема поставки заказа
type ShippingSystem struct{}

// Поставка заказа
func (s *ShippingSystem) ShipOrder(orderId string) error {
	fmt.Printf("Поставка заказа \"%s\"\n", orderId)
	return nil
}

// Фасад обработки заказов
type OrderProcessingFacade struct {
	inventory *InventorySystem // Подсистема инвентаря
	payment   *PaymentGateway  // Подсистема платежного шлюза
	shipping  *ShippingSystem  // Подсистема поставки
}

func NewOrderProcessingFacade() *OrderProcessingFacade {
	return &OrderProcessingFacade{
		inventory: &InventorySystem{},
		payment:   &PaymentGateway{},
		shipping:  &ShippingSystem{},
	}
}

// Обработка заказа
func (o *OrderProcessingFacade) ProcessOrder(productId string, quantity int, amount float64, orderId string) error {
	err := o.inventory.CheckInventory(productId, quantity)
	if err != nil {
		return fmt.Errorf("Не удалось проверить инвентарь: %v", err)
	}

	err = o.payment.ProcessPayment(amount)
	if err != nil {
		return fmt.Errorf("Не удалось обработать платеж: %v", err)
	}

	err = o.shipping.ShipOrder(orderId)
	if err != nil {
		return fmt.Errorf("Не удалось поставить заказ: %v", err)
	}

	return nil
}

func main() {
	facade := NewOrderProcessingFacade()
	err := facade.ProcessOrder("товар1", 6, 2500.0, "заказ1")
	if err != nil {
		log.Fatal("Не удалось обработать заказ: %v\n", err)
	}
}
