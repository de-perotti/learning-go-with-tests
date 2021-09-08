package adder_test

import (
	"fmt"
	"github.com/de-perotti/learning-go-with-tests/pkgs/adder"
	"github.com/de-perotti/learning-go-with-tests/pkgs/notates"
)

func ExampleAdder() {
	fmt.Printf("Sum is: %d", adder.Adder(2, 2))
	// Output: Sum is: 4
}

func ExampleRolezinho() {
	notates.Rolezinho()
	// Output: é nois
}
