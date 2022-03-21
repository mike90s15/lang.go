// https://youtu.be/VkuZ4QMnoZM 
package main

import (
    "fmt"
)

var x int = 42
var y string = "james bond"
var z bool = true

func main() {
    s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
    fmt.Print(s)
}
