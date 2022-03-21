//

package main

import "fmt"

func main() {
    list := []string{"item0", "item1", "item2", "item3", "item4"}
    for i := 0; i < len(list); i++ {
	fmt.Println("\033[1;31m" + list[i])
    }
}
