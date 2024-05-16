package main

import "fmt"

type Person struct {
	_name string
	_age  int
}

func (person *Person) init(name string, age int) {
	person._name = name
	person._age = age
}

func (person *Person) print() {
	fmt.Println(person._name, person._age)
}

func main() {
	p := Person{"Ahri", 13}
	p.print()
}
