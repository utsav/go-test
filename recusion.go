package main

import "fmt"

func main() {
	fact7 := factorial(7);
	fmt.Println(fact7);
}

func factorial(a int)int {
	if a == 0 {
		return 1;
	}
	return a * factorial(a - 1);
}