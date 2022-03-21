// 

package main

import "fmt"

func main() {
	a := h()
	fmt.Println(a)
}

func h() func() int {
    s := 0
    return func() int {
	s *= 10
	return s //fmt.Println(s)
    }
}

