package main

import (
	"fmt"

	puppy "github.com/joyadauche/test-dependency-one"
)

func main() {
	fmt.Println("Hello G😀!")

	a := puppy.Bark()
	b := puppy.Barks()
	fmt.Println(a)
	fmt.Println(b)

	c := puppy.BigBark()
	d := puppy.BigBarks()
	fmt.Println(c)
	fmt.Println(d)
}
