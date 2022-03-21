//

package main

import "fmt"

func main() {
    slice := []int{00, 01, 02, 03, 05, 06, 07, 8, 9}	
    fmt.Println(slice[:3])
    fmt.Println(slice[4:])
    fmt.Println(slice[1:7])
    fmt.Println(slice[2:9])
    fmt.Println(slice[2:len(slice) - 1])
}
