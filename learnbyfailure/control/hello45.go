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
	names := [2]string{"Jaehoon","Jaehyun"}
	ages := [2]int{20,21}

	for ( i := 0 ; i < len(names) ; i++ ) {
		hello(ages[i],names[i])
	}
}
