// https://youtu.be/5K17jFDXvWw 

package main

import (
   "fmt"
)

func main() {
    x := funcao()
    x()
}
    func funcao() func(){return func(){fmt.Println("done")}}

