package channeltest

import (
	"strconv"
	"time"
)

type Car struct {
	Val string
}

func MakeTire(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.Val += "Tire, "

		outChan <- car
	}
}

func MakeEngine(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.Val += "Engine, "

		outChan <- car
	}
}

func StartWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{Val: "Car" + strconv.Itoa(i)}
		i++
	}
}
