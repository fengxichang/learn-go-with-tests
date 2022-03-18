package structure

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//func TestArea(t *testing.T) {
//	checkArea := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{12.0, 6.0}
//		want := 72.0
//
//		checkArea(t, rectangle, want)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10.0}
//		want := 314.1592653589793
//
//		checkArea(t, circle, want)
//	})
//}

// 重构Area测试
func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{w: 12.0, h: 6.0}, want: 72.0},
		{name: "circle", shape: Circle{r: 10.0}, want: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
