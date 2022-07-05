package main

import (
	"fmt"
	"learngo/workshop"
)

func main() {
	fmt.Println("test")

	var name string
	fmt.Scanln(&name)

	a, _ := add(1, 2)
	fmt.Println(a)

	workshop.Print("something")

	for test := 0; test < 10; test++ {
		fmt.Printf(" I am number %v\n", test)
	}

	//var people [10] int{1,2,3,4,5,6,7,8,9,0}
	people := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i := 0; i < len(people); i++ {
		fmt.Printf(" I am person %v\n", people[i])
	}

	for _, person := range people {
		fmt.Printf(" I am person %v\n", person)
	}

	for i, _ := range people {
		fmt.Printf(" I am person %v\n", people[i])
	}

	var newIntern = intern{
		name:   name,
		gender: "M",
		age:    21,
	}

	fmt.Println(newIntern)
}

func add(a int, b int) (int, int) {
	return a + b, a - b
}

type intern struct {
	name   string
	gender string
	age    int
}

// a struct , get values with scanln
