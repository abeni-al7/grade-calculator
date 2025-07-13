package main

import "fmt"

var (
	name string
	amountOfSubjects int
	subjects = make(map[string]float64)
)

func acceptBasicInfo() {
	fmt.Println("Welcome to the Student Grade Tracking App")
	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Println("Hello,", name)
	fmt.Println("How many subjects have you taken?")
	fmt.Scanln(&amountOfSubjects)
}

func validateGrade(grade float64) bool {
	if grade >= 0 && grade <= 100 {
		return true
	}
	return false
}

func acceptSubject(i int) {
	var name string
	fmt.Println("Enter the name of Subject", i + 1)
	fmt.Scanln(&name)

	var grade float64
	fmt.Println("What is your grade for", name, "?")
	fmt.Scanln(&grade)

	for !validateGrade(grade) {
		fmt.Println("Please enter a valid grade between 0 and 100 (inclusive)")
		fmt.Scanln(&grade)
	}
	subjects[name] = grade
}

func calculateAverage() float64 {
	total := 0.0
	count := 0.0
	for _, value := range subjects {
		total += value
		count += 1.0
	}
	return total / count
}

func showOutput() {
	fmt.Printf("Hey, %v", name)
	fmt.Println()
	fmt.Println("Your scores for all your subjects are displayed below")
	for name, grade := range subjects {
		fmt.Printf("%v: %v", name, grade)
		fmt.Println()
	}
	fmt.Printf("As per the above results your average grade is %v", calculateAverage())
	fmt.Println()
}

func main() {
	acceptBasicInfo()

	i := 0
	for i < amountOfSubjects {
		acceptSubject(i)
		i++
	}
	showOutput()
}