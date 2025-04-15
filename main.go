package main

import (
	"fmt"

	"github.com/srtk1996/golang_practice_task/mystruct"
)

func main() {
	// Create an exported struct (Person) from mystruct
	q := mystruct.Person{Name: "Bob", Age: 30}
	fmt.Println("Name:", q.Name)
	fmt.Println("Age:", q.Age)
}
