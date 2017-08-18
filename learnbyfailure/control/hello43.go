package main

import "fmt"

func hello(age int,name string) {
	if  age < 15 {
		fmt.Printf("Hi, %v!\n",name);
	} else 
		fmt.Printf("Hello, %v!\n",name);
}

func main() {
	hello(20,"Jaehoon");
	hello(21,"Jaehyun");
}
