package main

import "fmt"

type Person struct {
	name string
	age int
}


func (person Person) PlusAge() {
      person.age++
}

func (person *Person) PlusAgePointer() {
      person.age++
}

func main() {
	person1 := Person{}
	person2 := Person{"Jaehoon",20}
	person3 := Person{name : "Jaehoon", age : 30}
	person4 := Person{age : 25, name : "Jaehoon"}
	var person5 *Person = new(Person)

	fmt.Println(person1);
	fmt.Println(person2);
	fmt.Println(person3);
	fmt.Println(person4);
	fmt.Println(person5);

	person1.PlusAge()
	person5.PlusAge()

	fmt.Println("Person1 %v",person1)
	fmt.Println("Person5 %v",person5)

	person1.PlusAgePointer()
	person5.PlusAgePointer()

	fmt.Println("Person1 %v",person1)
	fmt.Println("Person5 %v",person5)
}

