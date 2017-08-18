package main

import "fmt"

func hello(string name) {
	fmt.Println("Hello,",name,"!")
}

func main() {
	hello("Jaehoon");
	hello("Jaehyun");

	fmt.Println("End");
}
