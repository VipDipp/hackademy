package main

import (
	"fmt"
	"strconv"
	"time"
)

type Barber struct {
	state string
}

func NewBarber() (b *Barber) {
	return &Barber{
		state: "Sleeping",
	}
}

type Customer struct {
	name string
}

func barber(b *Barber, waitingRoom chan *Customer, arrived chan *Customer) {
	for {
		select {
		case c := <-waitingRoom:
			b.state = "Cutting"
			fmt.Printf(" %s ", c.name)
			fmt.Println("")
			time.Sleep(time.Millisecond * 100)

		default:
			select {
			case c := <-arrived:
				b.state = "Cutting"
				fmt.Printf(" %s ", c.name)
				fmt.Println("")
				time.Sleep(time.Millisecond * 100)
			default:
				time.Sleep(time.Millisecond * 100)
				fmt.Printf(".")
				b.state = "Sleeping"
			}
		}
	}
}

func customer(c *Customer, b *Barber, waitingRoom chan<- *Customer, arrived chan<- *Customer) {
	switch b.state {
	case "Sleeping":
		select {
		case arrived <- c:
			fmt.Println("WAKE UP")
		default:
			select {
			case waitingRoom <- c:
			default:
				return
			}
		}
	case "Cutting":
		select {
		case waitingRoom <- c:
			fmt.Println("New Customer")
		default:
			fmt.Println("Full")
			return
		}
	}
}

func main() {
	b := NewBarber()
	waitingRoom := make(chan *Customer, 5)
	arrived := make(chan *Customer, 1)
	go barber(b, waitingRoom, arrived)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		c := new(Customer)
		c.name = "Customer" + strconv.Itoa(i)
		go customer(c, b, waitingRoom, arrived)
	}
	time.Sleep(time.Millisecond * 1000)
	c := new(Customer)
	c.name = "late"
	go customer(c, b, waitingRoom, arrived)
	time.Sleep(time.Millisecond * 100000)
}
