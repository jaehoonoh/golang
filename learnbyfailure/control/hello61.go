package main

import "fmt"

func hello(age int,name string) {
	switch {
	case age < 15 :
		fmt.Printf("Hi, %v!\n",name);
	case age < 30 :
		fmt.Printf("Hello, %v!\n",name);
	default:
		fmt.Printf("Hello, %v! How are you?\n",name);
	}
}

func main() {
	hello(20,"Jaehoon");
	hello(21,"Jaehyun");
	hello(29,"goblinok");
	hello(30,"leeyeon");
}
