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
		// Сначала в общем виде говорим, что за объекты тут лежат
		shape Shape
		want  float64
	}{
		// После этого создаём слайс с этими объектами.
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
		// запятая в конце - это НЕ ошибка
	}
	for _, tt := range areaTest {
		got := tt.shape.Area()
		want := tt.want
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
}
