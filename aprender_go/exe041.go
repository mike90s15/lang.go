/*


*/
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 5, 7, 11}
    fmt.Println(funcao1(slice...))
    slice2 := []int{11, 21, 33, 51, 71, 11}
    fmt.Println(funcao2(slice2))
}

func funcao1(a ...int) int {
    v := 1
    for _, valor := range a {v *= valor}
    return v
}

func funcao2(a []int) int {
    v := 1
    for _, valor := range a {v *= valor}
    return v
}
