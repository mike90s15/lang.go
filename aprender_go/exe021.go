// https://youtu.be/5K17jFDXvWw 

package main

import (
   "fmt"
)

func main() {
    z := false
    b := true
    if z {
    	fmt.Println(z)
    } else if !b {
	fmt.Println(b)
    } else {
	fmt.Println("null")
   }
}
