package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Var Test")

	var name string = "Artur"
	var age = 18
	var version float32 = 0.1

	fmt.Println("\nHello, sr.", name, "you are", age)
	fmt.Println("\nProgram Version", version)

	fmt.Println("Var age type is", reflect.TypeOf(age))
}
