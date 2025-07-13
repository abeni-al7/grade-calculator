package main

import "fmt"

var (
	name string
	amountOfSubjects int
	subjects = make(map[string]float64)
)

func main() {

	fmt.Println("Welcome to the Student Grade Tracking App")
	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Println("Hello,", name)
	fmt.Println("How many subjects have you taken?")
	fmt.Scanln(&amountOfSubjects)

	i := 0
	for i < amountOfSubjects {
		var name string
		fmt.Println("Enter the name of Subject", i)
		fmt.Scanln(&name)

		var grade float64
		fmt.Println("What is your grade for", name, "?")
		fmt.Scanln(&grade)

		subjects[name] = grade
		i++
	}
	fmt.Println(subjects)
}