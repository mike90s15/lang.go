//

package main

import "fmt"

func main() {
     mape := map[string][]string{
	"Mike Edwards": []string{
            "Programa", "aritmética", "viajar",
            },
	"Pessoa2": []string{
            "P2 item1", "P2 item2", "P2 item3",
	    },
	"Pessoa3": []string{
	    "P3 item1", "P3 item2", "P3 item3",	
	    },
	}

    for k, v := range mape {
	fmt.Println(k)
	for i, h := range v {
            fmt.Println("\t", i, " - ", h)
	}
    }
}

