package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Barber struct {
	state      string
	isSleeping bool
}

type BarberShop struct {
	waitingRoom chan *Customer
	arrived     chan *Customer
}

func NewBarber() (b *Barber) {
	return &Barber{
		state:      "Sleeping",
		isSleeping: false,
	}
}

func NewBarberShop(lobbySpace int, arriveSpace int) (bs *BarberShop) {
	return &BarberShop{
		waitingRoom: make(chan *Customer, lobbySpace),
		arrived:     make(chan *Customer, arriveSpace),
	}
}

type Customer struct {
	name string
}

func customerFlow(amount int, b *Barber, bs *BarberShop) {
	names := []string{
		"Bob",
		"Sam",
		"Danylo",
		"Sasha",
		"Matviy",
		"Ivan",
		"Jack",
		"Lucifer",
		"Arthas",
		"Uter",
	}
	for i := 0; i < amount; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)*10))
		c := new(Customer)
		c.name = names[rand.Intn(10)]
		go customerEnter(c, b, bs)
	}
}

func barberWork(b *Barber, bs *BarberShop) {
	for {
		select {
		case c := <-bs.waitingRoom:
			b.state = "Cutting"
			fmt.Println("	Barber is cutting " + c.name)
			time.Sleep(time.Millisecond * 100)
			fmt.Println("	Barber done cutting " + c.name)
		case c := <-bs.arrived:
			b.state = "Cutting"
			fmt.Println("	Barber is cutting " + c.name)
			time.Sleep(time.Millisecond * 100)
			fmt.Println("	Barber done cutting " + c.name)
		default:
			if !b.isSleeping {
				fmt.Println("		Barber go to sleep")
				b.isSleeping = true
			}
			b.state = "Sleeping"
		}
	}
}

func customerEnter(c *Customer, b *Barber, bs *BarberShop) {
	switch b.state {
	case "Sleeping":
		select {
		case bs.arrived <- c:
			fmt.Println(c.name + " came, barber wake up")
			b.isSleeping = false
		}
	case "Cutting":
		select {
		case bs.waitingRoom <- c:
			fmt.Println(c.name + " is waiting in lobby")
		default:
			fmt.Println("		Lobby is full " + c.name + " is leaving but promised to return")
			time.Sleep(time.Millisecond * 2000)
			go customerEnter(c, b, bs)
		}
	}
}

func main() {
	barber := NewBarber()
	barberShop := NewBarberShop(5, 1)
	go barberWork(barber, barberShop)
	customerFlow(10, barber, barberShop)

	time.Sleep(time.Millisecond * 1000)
	customerFlow(1, barber, barberShop)

	for {

	}
}

/*
	to run you should initialize barber, create barbershop, and customer flow with amount of customers coming in barbershop,
	also you can send one customer as well - initialize customer with name and run customerEnter with customer barbershop and barber
*/
