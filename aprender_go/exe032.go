//

package main

import "fmt"

func main() {
    mother := [][]string{
		[]string{
		  "Ana",
		  "lima",
		  "dormir"},
		[]string{
		  "mimi",
		   "<3",
		   "ser linda"},
		[]string{
		  "meus queridos",
		   "que estudam bastante",
		   "fazer os exercícios ninja"}}
    for _, v := range mother {
	for _, item := range v {
	    fmt.Println(item)
	}
	fmt.Println()

    }
}

