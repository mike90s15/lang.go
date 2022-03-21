// 

package main

import "fmt"

func main() {
    esporte := "Parkour"
    switch esporte {
        case "futebol":
	    fmt.Println("case x == 1:")
	case "corrida":
	    fmt.Println("case x == 2:")
	case "Parkour":
	    fmt.Println("You love Parkour")
    }
}			
