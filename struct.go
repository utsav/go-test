package main

import "fmt"

type Person struct {
	name string; age int
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(&Person{name: "Ann", age: 40})

	s := Person{name: "Sean", age: 50}

	sp := &s
	fmt.Println(sp.age)

	sp.age = 60

	fmt.Println("sp.age ", sp.age)
	fmt.Println("s.age ", s.age)
}
