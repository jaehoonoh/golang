package main

import "fmt"

func test() {
	var make int

	fmt.Println(make);
}

func main() {
	var ages []int = make([]int,2,3)
	var heights []int = make([]int,5,10)

	ages[0] = 1
	ages[1] = 2

	x := append(ages,3)

	fmt.Println(ages);
	fmt.Println(x);

	copy(ages,heights)
	fmt.Println(ages);
	fmt.Println(heights);

	copy(heights,ages)
	fmt.Println(ages);
	fmt.Println(heights);

	test()
}

