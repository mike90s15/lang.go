//


package main

import (
    "fmt"
)

func main() {
    p := 300
    fmt.Printf("%d\t%b\t%#X\n", p, p, p)
    r := p << 4
    fmt.Printf("%d\t%b\t%#X\n", r, r, r)

}
