package structs

// A struct is a sequence of named elements, called fields, each of which has a name and a type.
// Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).
type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}