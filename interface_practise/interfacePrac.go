package main 

import "fmt"
import "math"

type circle struct{
	radius float64
}

type square struct{
	side float64
}

func (c circle) area() float64{
	return math.Pi*math.Pow((c.radius),2)
}

func (s square) area() float64{
	return math.Pow(s.side,2)
}

type shape interface{
	area() float64
}

func Printer(s shape){
	fmt.Println("The are is ",s.area())
}

func main(){
	x:=square{}
	x.side=10
	Printer(x)
	circle:=circle{}
	circle.radius=math.Pi
	Printer(circle)
}