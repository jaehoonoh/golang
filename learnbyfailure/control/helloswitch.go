package main

import "fmt"

func hello(age int,name string) {
	switch age {
	case 0: fallthrough
	case 2,4,68:
		fmt.Printf("Hi, %v!\n",name);
	case 1,3,5,7,9:
		fmt.Printf("Hello, %v!\n",name);
	default:
		fmt.Printf("Hello, %v! How are you?\n",name);
	}
}

func main() {
	hello(0,"Baby");
	hello(1,"Baby2");
	hello(20,"Jaehoon");
	hello(21,"Jaehyun");
	hello(29,"goblinok");
	hello(30,"leeyeon");
}
