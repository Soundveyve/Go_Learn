package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Первый табличный тест
func TestArea(t *testing.T) {
	// Shape - интерфейс, который должны реализовывать Circle() и Reactangle()
	areaTest := []struct {
		name string
		// Сначала в общем виде говорим, что за объекты тут лежат
		shape   Shape
		hasArea float64
	}{
		// После этого создаём слайс с этими объектами.
		{name: "Прямоугольник", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Круг", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Треугольник", shape: Triangle{12, 6}, hasArea: 36.0},
		// запятая в конце - это НЕ ошибка
	}
	for _, tt := range areaTest {
		got := tt.shape.Area()
		want := tt.hasArea
		if got != want {
			t.Errorf("%#v got %g want %g", tt.shape, got, want)
		}
	}
}
