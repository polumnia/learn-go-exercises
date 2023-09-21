package structs

import "math"

// An interface type defines a type set.
// A variable of interface type can store a value of any type that is in the type set of the interface.
// Such a type is said to implement the interface.
// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

type Shape interface {
	Area() float64
}

// A struct is a sequence of named elements, called fields, each of which has a name and a type.
// Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).
type Rectangle struct {
	Width float64
	Height float64
}

// method syntax:
// func (receiverName ReceiverType) MethodName(args)

// method of Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// method of Circle struct
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}