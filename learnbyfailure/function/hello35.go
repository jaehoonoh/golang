package main

import "fmt"

func hello(name string) (result string) {
	result = "Hello, " + name + "!"
}

func main() {
	fmt.Println(hello("Jaehoon"));
	fmt.Println(hello("Jaehyun"));
}
