// https://youtu.be/5K17jFDXvWw 

package main

import (
   "fmt"
)

func main() {
    for x := 65; x <= 90; x++ {
        fmt.Println()
	for i := 1; i <= 3; i++ {
	    fmt.Printf("\t%#U", x)	
	}
    }
    fmt.Println()
}
