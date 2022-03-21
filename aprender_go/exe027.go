//

package main

import "fmt"

func main() {
    slice := []int{00, 01, 02, 03, 05, 06, 07, 8, 9}
 
    for i, v := range slice {
   	fmt.Println(i, "Valor", v)	
    }
    fmt.Printf("\033[33m%T\n", slice)
}
