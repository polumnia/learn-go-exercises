package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// The f is for our float64 and the .2 means print 2 decimal places.
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// example of table driven tests
func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		// The test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations
		{name: "Rectangle", shape: Rectangle{5.0, 10.0}, hasArea: 50.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
		
	}
}