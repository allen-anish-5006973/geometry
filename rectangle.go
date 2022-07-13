package geometry

type Rectangle struct {
	length  float64
	breadth float64
}

func NewRectangle(length, breadth float64) Rectangle {
	if length <= 0 || breadth <= 0 {
		panic("length and breadth must be positive")
	}
	return Rectangle{length, breadth}
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}
