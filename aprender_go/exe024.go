//

package main

import "fmt"

func main() {
     fmt.Println(true && true)   // True
     fmt.Println(true && false)  // False
     fmt.Println(false && true)  // False
     fmt.Println(true || true)   // True
     fmt.Println(true || false)  // True
     fmt.Println(false || true)  // True  
     fmt.Println(!true)		 // False
     fmt.Println(!false)	 // True
}
