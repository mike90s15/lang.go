// 

package main

import "fmt"

type pessoa struct {nome string; sobrenome string; idade int}
	
func (x pessoa) demostra() {fmt.Println(x.nome, x.sobrenome, x.idade)}

func main() {
    xx := pessoa {
	    nome: "Mike",
	    sobrenome: "edwards",
	    idade: 23,
	}
    bb := pessoa{nome: "sla", sobrenome: "não tem", idade: 0}
    bb.demostra()
    xx.demostra()
}
