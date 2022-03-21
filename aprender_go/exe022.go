// 

package main

import "fmt"

func main() {
    x := 2
    switch {
        case x == 1:
	    fmt.Println("case x == 1:")
	case x == 2:
	    fmt.Println("case x == 2:")
	case x == 3:
	    fmt.Println("case x == 3:")
    }
}
