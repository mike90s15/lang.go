//

package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
        sabores []string

}

func main() {
	pessoa1 := pessoa {
		nome: "Renata",
		sobrenome: "Pimenta",
		sabores: []string{"Chocolate", "Morango", "Baunilha"}}
	pessoa2 := pessoa {"frederico", "da Prússia", 
		[]string{"Sabão em pó", "pé de coelho", "feijão"}}
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	
	for _, va := range pessoa1.sabores {fmt.Println("\t", va)}
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, va := range pessoa1.sabores {fmt.Println("\t", va)}
}

