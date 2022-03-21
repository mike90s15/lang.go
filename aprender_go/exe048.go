// https://youtu.be/5K17jFDXvWw 

package main

import (
   "fmt"
)

func main() {
    funcao2(funcao)
}
    func funcao(){fmt.Println("sla1")}
    func funcao2(x func()){fmt.Println(x);fmt.Println("sla2")}

