package main

import (
	"fmt"
)

func main() {
	fmt.Println("Circle------>\n")
	var c Circle = Circle{12, 54, 12} //first way how to init var
	fmt.Println(c)

	c1 := Circle{100, 100, 100} //second way how to init var
	fmt.Println(c1)

	c2 := Circle{x: 100, y: 100, z: 100} //third way how to init var
	fmt.Println(c2)

	circle := new(Circle) //get pointer in structure, all fields will get default values
	fmt.Println(circle)
	getField(&c2)

	fmt.Println("\n\nRectagle ------>\n")
	r := Rectangle{5, 10}
	fmt.Printf("Area of rectangle with a = %0.f, b = %0.f is S = %0.f", r.x, r.y, r.average())

	fmt.Println("\n\nAnonymous types ---->\n")
	boy := Boy{Apple{100}, "Micheal"}
	fmt.Printf("%s has %d apple(s)\n", boy.Name, boy.PrintCount())

	boy2 := new(Boy)
	boy2.count = 3
	boy2.Name = "John"
	fmt.Printf("%s has %d apple(s)", boy2.Name, boy2.PrintCount())

	shape1 := Circle{50, 50, 50}
	shape2 := Rectangle{5, 10}
	fmt.Println("\n\nShapes ------>\n")
	fmt.Println(totalAverage(&shape1, &shape2))

	fmt.Println("\n\nMultiShapes ------>\n")

	shapes := make([]Shape, 2)
	shapes[0] = &shape1
	shapes[1] = &shape2

	ms := MultiShape{shapes}

	fmt.Println("Average of multishape", ms.average())

}
// region our types
type Circle struct {
	x, y, z int
}

type Rectangle struct {
	x, y float64
}

type Boy struct {
	// Apple field call an anonymous type, if we will write as it, we can get access to Apples method PrintCount(),
	// otherwise we should write boy.Apple.PrintCount(), but we also could write as it even if we have anonymous type
	Apple
	Name string
}
type Apple struct {
	count int
}

type Shape interface {
	average() float64

}
type MultiShape struct {
	shapes []Shape
}

//endregion

func getField(circle *Circle) {
	//this is a simple function
	fmt.Printf("Fields: x = %d, y = %d, z = %d ", circle.x, circle.y, circle.z)
}

func (c *Circle) average() float64 {
	//this is a method of Circle structure
	return (float64)(c.x + c.y + c.z) / 3;
}

func (r *Rectangle) average() float64 {
	// this is a method of Rectangle structure
	return (r.x + r.y) / 2;
}

func (a *Apple) PrintCount() int {
	return a.count;
}

func totalAverage(shapes ...Shape) (total float64) {
	for _, s := range shapes {
		total += s.average()
	}
	return
}

func (m *MultiShape) average() (average float64) {
	for _, shape := range m.shapes {
		average += shape.average()
	}
	return average / float64(len(m.shapes))
}

