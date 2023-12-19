package c4areapoligon

import "fmt"

type Triangle struct {
	h float64
	b float64
}

type Square struct {
	a float64
}

type Rectangle struct {
	a float64
	b float64
}

func (t Triangle) area() float64 {
	return t.b * t.h / 2
}
func (s Square) area() float64 {
	return s.a * s.a
}
func (r Rectangle) area() float64 {
	return r.a * r.b
}

type AreaPoligonos interface {
	area() float64
}

func Areapol(shape AreaPoligonos) {
	fmt.Println("*******Area of Poligon***********")
	fmt.Printf("%T %v:\n", shape, shape.area())

}

// func main() {
// 	t := triangle{2.3, 4.4}
// 	q := square{3.3}
// 	r := rectangle{3.3, 6.6}
// 	Areapol(t)
// 	Areapol(q)
// 	Areapol(r)
// }
