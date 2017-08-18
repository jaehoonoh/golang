package main

import "fmt"

func hello(age int,name string) {
	if ( age < 15 ) {
		fmt.Printf("Hi, %v!\n",name);
	} else {
		fmt.Printf("Hello, %v!\n",name);
	}
}

func main() {
	for i,j := 0,0 ; i < 10 ; i+=1, j+=2  {
		fmt.Println(i," ",j);
	}
}
