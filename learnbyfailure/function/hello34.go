package main

import "fmt"

func hello(name string) result string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(hello("Jaehoon"));
	fmt.Println(hello("Jaehyun"));
}
