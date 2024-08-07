package structural

import "fmt"

type Shape interface {
	Draw(x, y int)
}

type Circle struct {
	color string
}

func (c *Circle) Draw(x, y int) {
	fmt.Printf("Drawing Circle [color: %s, x: %d, y: %d]\n", c.color, x, y)
}

func NewCircle(color string) *Circle {
	return &Circle{color: color}
}

// Rectangle struct
type Rectangle struct {
	color string
}

func (r *Rectangle) Draw(x, y int) {
	fmt.Printf("Drawing Rectangle [color: %s, x: %d, y: %d]\n", r.color, x, y)
}

func NewRectangle(color string) *Rectangle {
	return &Rectangle{color: color}
}

type ShapeFactory struct {
	circleMap    map[string]*Circle
	rectangleMap map[string]*Rectangle
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap:    make(map[string]*Circle),
		rectangleMap: make(map[string]*Rectangle),
	}
}

func (s *ShapeFactory) GetCircle(color string) *Circle {
	if circle, exists := s.circleMap[color]; exists {
		return circle
	}
	circle := NewCircle(color)
	s.circleMap[color] = circle
	return circle
}

func (s *ShapeFactory) GetRectangle(color string) *Rectangle {
	if rectangle, exists := s.rectangleMap[color]; exists {
		return rectangle
	}
	rectangle := NewRectangle(color)
	s.rectangleMap[color] = rectangle
	return rectangle
}

func runCircle() {
	shapeFactory := NewShapeFactory()

	circle1 := shapeFactory.GetCircle("red")
	circle1.Draw(10, 20)

	circle2 := shapeFactory.GetCircle("Red")
	circle2.Draw(30, 40)

	circle3 := shapeFactory.GetCircle("Blue")
	circle3.Draw(50, 60)

}
