// 

package main

import ("fmt"; "math")

type quadrado struct{lado1 float64; lado2 float64}
type circulo struct{raio float64}
type info interface{area()}

func (x quadrado) area(){fmt.Println("Area é", x.lado1 * x.lado2)}
func (x circulo) area(){fmt.Println("Area é", math.Pi * 2 * x.raio)}
func medida(i info){i.area()}

func main() {
    xx := quadrado{lado1: 10, lado2: 2}
    bb := circulo{raio: 42}
    medida(bb)
    medida(xx)
}
