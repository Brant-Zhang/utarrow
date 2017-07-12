package main

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() CarModel
}

type CarModel interface {
	Drive() error
	Stop() error
}

type CarAttr struct {
	color Color
	wheel Wheels
	speed Speed
}

func (this *CarAttr) Drive() error {
	fmt.Printf("this car is driving in speed:%.2f\n", this.speed)
	return nil
}

func (this *CarAttr) Stop() error {
	fmt.Println("this car is stoped!")
	return nil
}

func (this *CarAttr) Color(c Color) Builder {
	this.color = c
	return this
}

func (this *CarAttr) Wheels(w Wheels) Builder {
	this.wheel = w
	return this
}

func (this *CarAttr) TopSpeed(s Speed) Builder {
	this.speed = s
	return this
}

func (this *CarAttr) Build() CarModel {
	return this
}

func NewBuilder() *CarAttr {
	return new(CarAttr)
}

func main() {
	sportCar := NewBuilder().Color(BlueColor).Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	sportCar.Drive()
	familyCar := NewBuilder().Color(RedColor).Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	familyCar.Drive()
}
