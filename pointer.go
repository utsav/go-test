package main

import "fmt"

func main() {
	i := 0;

	fmt.Println("i ", i);
	withoutRef(i)
	fmt.Println("i ", i);
	withRef(&i);
	fmt.Println("i ", i);
}

func withoutRef(ival int) {
	ival = 1;
}

func withRef(iptr *int) {
	*iptr = 1;
}
