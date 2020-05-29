package main

import (
	"fmt"
	"math"
)

type Point struct{
	x,y float64
}

func (p *Point) Abs() float64{
	return math.Sqrt(p.x*p.x+p.y*p.y)
}

type NamedPoint struct{
	a Point
	name string
}

func main(){
	n:=&NamedPoint
	*n.a={3,4}
	*n.name="sedg"
	fmt.Println(n.Abs())
}
