package main

import "fmt"

func main() {
	var minimumWage int = 10
	var workingHour int = 20

	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, minimumWage * workingHour, income)
	// 10 20 200 200
}