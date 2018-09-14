package main

import "fmt"

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
        fmt.Println(msg)
    }("going")
	fmt.Scanln()
    fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}
