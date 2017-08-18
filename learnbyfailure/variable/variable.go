package main

import "fmt"

func main() {
	var age int
	var height int = 200
	weight := 80

	var number1,number2 int
	var number3,number4 int = 3,4
	var number5 ,number6 = 5,6
	var number7 ,name = 7,"Jaehoon"
	number8 ,name2 := 7,"Jaehoon"

	fmt.Println("Age =",age)
	fmt.Println("Height =",height)
	fmt.Println("Weight =",weight)

	fmt.Printf("%v %v\n",number1,number2)
	fmt.Printf("%v %v\n",number3,number4)
	fmt.Printf("%v %v\n",number5,number6)
	fmt.Printf("%v %v\n",number7,name)
	fmt.Printf("%v %v\n",number8,name2)
}

