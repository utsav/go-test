package main

import "fmt"

func main(){
	hello()
	var num = 9;
	num = 10;
	fmt.Println("num is ", num);
	var e string;
	fmt.Println(len(e))
	fmt.Println("1 + 1 =", 1.0 + 1.0)
	fmt.Println("Hello World"[1])

}

func hello(){
	var num = 25;
	fmt.Println("num is ", num)
}
