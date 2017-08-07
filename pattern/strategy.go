package pattern

import (
	"fmt"
)

//strategy ,defined an common interface
type FlyBehavior interface {
	fly()
}

//concrete strategy
//each concrete strategy implements an algorithm
type FlyWithWings struct {
}

func (f *FlyWithWings) fly() {
	fmt.Println("i can fly in the sky")
}

type FlyDisable struct{}

func (f *FlyDisable) fly() {
	fmt.Println("i cannot fly")
}

//strategy
type QuackBehavior interface {
	quack()
}

//concrete strategy
type QuackCommon struct{}

func (q *QuackCommon) quack() {
	fmt.Println("the duck is quacking")
}

type QuackSilent struct{}

func (q *QuackSilent) quack() {
	fmt.Println("I cann't speak!")
}

type Duck struct {
	Fly   FlyBehavior
	Quack QuackBehavior
	color string
	name  string
}

func (d *Duck) PerformFly() {
	d.Fly.fly()
}

func (d *Duck) PerformQuack() {
	d.Quack.quack()
}

func (d *Duck) Display() {
	fmt.Printf("the duck:%s is in %s color\n", d.name, d.color)
}

func NewMarkDuck() *Duck {
	d := new(Duck)
	f := new(FlyWithWings)
	q := new(QuackSilent)
	d.Fly = f
	d.Quack = q
	d.color = "black"
	d.name = "Screamer"
	return d
}

/*
func main() {
	d := NewMarkDuck()
	d.Display()
	d.PerformFly()
	d.PerformQuack()
}*/
