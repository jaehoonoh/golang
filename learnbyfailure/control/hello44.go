package main

import "fmt"

func hello(age int,name string) {
	if  age < 15 {
		fmt.Printf("Hi, %v!\n",name);
	} else if ( age < 30 )
	{
		fmt.Printf("Hello, %v!\n",name);
	} else {
		fmt.Printf("Hello, %v! How are you?\n",name);
	}
}

func main() {
	hello(20,"Jaehoon");
	hello(21,"Jaehyun");
	hello(29,"goblinok");
	hello(30,"leeyeon");
}
