package main

import "fmt"

// Интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Конкретная стратегия: оплата кредитной картой
type CardPayment struct {
	cardNumber string
}

func NewCardPayment(cardNumber string) *CardPayment {
	return &CardPayment{cardNumber}
}

func (c *CardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено %.2f р. с помощью банковской карты %s", amount, c.cardNumber)
}

// Конкретная стратегия: оплата через ЮMoney
type YooMoneyPayment struct {
	email string
}

func NewYooMoneyPayment(email string) *YooMoneyPayment {
	return &YooMoneyPayment{email}
}

func (p *YooMoneyPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено %.2f р. через ЮMoney с адреса %s", amount, p.email)
}

// Контекст, использующий стратегию оплаты
type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (s *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	s.paymentStrategy = strategy
}

func (s *ShoppingCart) Checkout(amount float64) string {
	return s.paymentStrategy.Pay(amount)
}

func main() {
	cart := NewShoppingCart()

	// Выбираем стратегию оплаты: банковская карта
	cardPayment := NewCardPayment("1234-5678-9876-5432")
	cart.SetPaymentStrategy(cardPayment)

	// Выполняем оплату
	result := cart.Checkout(2500.0)
	fmt.Println(result)

	// Меняем стратегию оплаты: ЮMoney
	yooPayment := NewYooMoneyPayment("example@example.com")
	cart.SetPaymentStrategy(yooPayment)

	// Выполняем оплату
	result = cart.Checkout(5000.0)
	fmt.Println(result)
}
