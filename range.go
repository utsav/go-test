package main

import "fmt"

func main() {
	nums := [3]string{"one", "two", "three"}

	for _, num := range nums {
		fmt.Println(num)
	}
}