package main

import "fmt"

func main(){
	x := []string{"utsav","bhavsar"}
	x= append(x, "j")
	for i, val := range x {
		fmt.Println(i, val);
	}
}
