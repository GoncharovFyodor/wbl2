package main

import "fmt"

type Department interface {
	execute(*Patient)
	// Передача в следующее отделение
	setNext(Department)
}

// Приемный покой
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Пациент уже зарегистрирован")
		r.next.execute(p)
		return
	}
	fmt.Println("Регистрация пациента")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

// Врач
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Осмотр у врача уже выполнен")
		d.next.execute(p)
		return
	}
	fmt.Println("Осмотр пациента врачом")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

// Медицинский кабинет
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Лекарство уже выдано")
		m.next.execute(p)
		return
	}
	fmt.Println("Выдача лекарства в медицинском кабинете")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

// Касса
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Оплата уже произведена")
	}
	fmt.Println("Производится оплата")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

// Пациент
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "фывапролдж"}

	reception.execute(patient)
}
