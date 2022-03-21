/*


*/
package main

import "fmt"

func main() {
    fmt.Println(ret_int())
    fmt.Println(ret_str())
}

func ret_int() int {
    return 0
}

func ret_str() (int, string) {
    return 10, "dez"
}
