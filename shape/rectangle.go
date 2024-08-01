package shape

type Rectangle struct {
	Length float32
	Width  float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.Length * rectangle.Width
}

func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Length + rectangle.Width)
}
