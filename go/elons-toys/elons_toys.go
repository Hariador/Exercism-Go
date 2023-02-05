package elon

import (
	"fmt"
)

func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance = car.distance + car.speed
		car.battery = car.battery - car.batteryDrain
	}

}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
	gos := trackDistance/car.speed + trackDistance%car.speed
	if car.battery >= gos*car.batteryDrain {
		return true
	}
	return false
}
