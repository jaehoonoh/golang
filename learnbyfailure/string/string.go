package main

import "fmt"
import "unicode/utf8"

func main() {
	name := "Jaehoon"
	var korean string = "재훈"
	var korean2 rune = '재'

	fmt.Println("Hello, ",name,korean, korean2, "!");
	fmt.Println("Length = ",len(name),utf8.RuneCountInString(korean));
}
