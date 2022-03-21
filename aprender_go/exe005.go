// https://youtu.be/O0r318FN_Uw

package main

import (
    "fmt"
)

type hotdog int
var x hotdog
var y int

func main() {
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    x = 42
    fmt.Println(x)
    y = int(x)
    fmt.Printf("↑ X e Y ↓\n%v\n%T\n", y, y)
}
