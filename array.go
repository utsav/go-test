package main

import "fmt"
import "encoding/json"

func main() {

	type Message struct {
		Name string
		Age int
		Time int64
	}
	var m Message
	b := []byte(`{"Name":"Utsav","Age":24,"Time":1294706395881547000}`)
	json.Unmarshal(b, &m)
	var a [5]byte
	a[1] = 1
	a[4] = 4
	fmt.Println("set 1", a)
	fmt.Println("set 2", b)
	fmt.Println("set 3", m.Name)
}

