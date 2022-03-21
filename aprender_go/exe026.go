//

package main

import "fmt"

func main() {
    array := [5]int{00, 01, 02, 03, 05}
 
    for i, v := range array {
   	fmt.Println(i, "Valor", v)	
    }
    fmt.Printf("\033[33m%T\n", array)
}
