package pattern

import "fmt"

type GrapheEditor struct {
}

func (g *GrapheEditor) drawColor(color string) {
	fmt.Println("this shape contains color:", color)
}

func (g *GrapheEditor) drawShape(s Shape) {
	s.draw()
}

type Shape interface {
	draw()
}

type Rectangle struct {
	x, y int
}

func (r *Rectangle) draw() {
	fmt.Println("draw rectangle:", r.x, r.y)
}

type Circle struct {
	diameter int
}

func (c *Circle) draw() {
	fmt.Println("draw circle:", c.diameter)
}
