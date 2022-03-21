package main

import ("fmt"; "sync")

var wg sync.WaitGroup

func main() {
    novasgoroutines(10)
    novasgoroutines(10)
    novasgoroutines(10)
    novasgoroutines(10)
    novasgoroutines(10)
    novasgoroutines(10)
    wg.Wait()
}

func novasgoroutines(i int) {
    wg.Add(i)
    for j := 0; j < i; j++ {
	x := j
	go func(i int) {
	fmt.Println("Eu sou a goroutine número:", i+1, "!")
	wg.Done()
	}(x)
    }
}
