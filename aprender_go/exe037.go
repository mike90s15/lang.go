//

package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
        sabores []string

}

func main() {
	meumape := make(map[string]pessoa)
	meumape["pimentão"]= pessoa {
		nome: "Renata",
		sobrenome: "Pimentão",
		sabores: []string{"Chocolate", "Morango", "Baunilha"}}
	meumape["da Prússia"] = pessoa {"frederico", "da Prússia", 
		[]string{"Sabão em pó", "pé de coelho", "feijão"}}
//	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, va := range meumape {
		fmt.Println("Meu nome é", va.nome, "meu sorvete favoritos são:")
		for _, va := range va.sabores {
			fmt.Println("-", va)
			
		}
		fmt.Println()
	}
//	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

//	for _, va := range meumape1 {fmt.Println("\t", va)}
	
//	for _, va := range meumape {fmt.Println("\t", va)}
}

