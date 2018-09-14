package main

import "fmt"
import "strconv"

func main() {

	res := plus("2", 4);
	a, b := vals();

	fmt.Println("res ", res);
	fmt.Println("a ", a);
	fmt.Println("b ", b);

}

func plus(a string, b int)string {
	return a + strconv.Itoa(b)
}

func vals() (int, int) {
	return 10, 20
}