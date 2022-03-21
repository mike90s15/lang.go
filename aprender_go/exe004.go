// https://youtu.be/5-0S-gefNe0

package main

import (
    "fmt"
)

type hotdog int
var x hotdog

func main() {
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    x = 42
    fmt.Println(x)
}
