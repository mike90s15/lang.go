// https://youtu.be/5K17jFDXvWw 

package main

import (
   "fmt"
)

func main() {
    for x := 1988; x <= 2089; x++ {
        fmt.Println("Ano: ", x)
	if ( x == 2008 ) {
	    break
	}
    }
}
