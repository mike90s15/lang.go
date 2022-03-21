//

package main

import "fmt"

func main() {
    x := `#include<stdio.h>
           
int main()
{
    printf("hello world\n");
    return 0;
}`
    
    fmt.Println(x)
}
