package structsmethodsinterfaces

import (
	"math"
)

type Shape interface {
	Area() float64
}

// Доступ до структур и их полей также как и функций определяется через верхний/нижний регистр
type Rectangle struct {
	Width  float64
	Height float64
}

/*
Дополнительный аргумент между названием и ключевым словомо func() - получатель.
Он нужен для того, чтобы привязать метод к опеределённой структуре
*/
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
