package area

import "testing"

func TestArea(t *testing.T) {
	testArea := []struct {
		name     string
		form     Form
		expected float64
	}{
		{name: "Rectangle", form: Rectangle{12, 6}, expected: 72.0},
		{name: "Circle", form: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", form: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range testArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.form.Area()
			if result != tt.expected {
				t.Errorf("%#v Result %.2f, expected %.2f", tt.form, result, tt.expected)
			}
		})
	}
}
