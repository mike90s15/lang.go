// https://youtu.be/5K17jFDXvWw 

package main

import (
   "fmt"
)

func main() {
    for x := 8208; x <= 12000; x++ {
	fmt.Printf("Caracte: %v - %v\n", x, string(x))
    }
}
