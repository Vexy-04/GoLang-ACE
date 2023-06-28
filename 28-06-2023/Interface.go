package main

import(
	"fmt"
	"math"
)

type Shape Interface{
	Area() float64
}

type Circle struct{
	radius float64
}

func (c Circle)Area()float64{
	return  math.Pi*c.radius*c.radius
}

type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle)Area()float64{
	return r.width*r.height
}

func main(){
	Circle := Circle{radius: 5}
	Rectangle := Rectangle{width: 4, height: 6}

	PrintArea(Circle)
	PrintArea(Rectangle)
}