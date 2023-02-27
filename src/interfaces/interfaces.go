package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (m Car) String() string {
	return fmt.Sprintf("Car: %v", string(m))
}

func (m Car) PickLift() Lift {
	return StandardLift
}

func (m Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(m))
}

func (m Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("Send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("Send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("Send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("Mountain crusher!")
	motorcycle := Motorcycle("Babeta")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
