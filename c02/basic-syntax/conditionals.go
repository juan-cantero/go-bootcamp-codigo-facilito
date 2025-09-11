package main

import "fmt"

// IfStatementExamples demonstrates various if statement patterns
func IfStatementExamples() {
	fmt.Println("=== IF STATEMENT EXAMPLES ===")
	
	// Simple if
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	}
	
	// If-else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
	
	// If with initialization
	if num := 42; num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
	
	// Multiple conditions
	temperature := 25
	humidity := 60
	if temperature > 20 && humidity < 70 {
		fmt.Println("Perfect weather!")
	}
	
	// Checking for zero values
	var name string
	if name == "" {
		fmt.Println("Name is empty")
	}
	
	// Checking slices
	numbers := []int{1, 2, 3}
	if len(numbers) > 0 {
		fmt.Printf("First number: %d\n", numbers[0])
	}
	
	fmt.Println()
}

// SwitchStatementExamples demonstrates various switch statement patterns
func SwitchStatementExamples() {
	fmt.Println("=== SWITCH STATEMENT EXAMPLES ===")
	
	// Basic switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend!")
	}
	
	// Switch with multiple values in case
	grade := 'B'
	switch grade {
	case 'A', 'B':
		fmt.Println("Excellent!")
	case 'C', 'D':
		fmt.Println("Good job!")
	case 'F':
		fmt.Println("Try harder!")
	default:
		fmt.Println("Invalid grade")
	}
	
	// Switch with conditions (no expression)
	temperature := 30
	switch {
	case temperature < 0:
		fmt.Println("Freezing!")
	case temperature >= 0 && temperature < 20:
		fmt.Println("Cold")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Comfortable")
	case temperature >= 30:
		fmt.Println("Hot!")
	}
	
	// Switch with initialization
	switch result := Calculate(10, 5); result {
	case 15:
		fmt.Println("Addition result")
	case 50:
		fmt.Println("Multiplication result")
	case 5:
		fmt.Println("Subtraction result")
	case 2:
		fmt.Println("Division result")
	default:
		fmt.Printf("Unknown operation result: %d\n", result)
	}
	
	// Switch on type (type switch)
	var value interface{} = "hello"
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
	
	// Switch with fallthrough (uncommon in Go)
	number := 1
	switch number {
	case 1:
		fmt.Println("One")
		fallthrough // Continue to next case
	case 2:
		fmt.Println("One or Two")
	case 3:
		fmt.Println("Three")
	}
	
	fmt.Println()
}

// Calculate is a helper function for switch example
func Calculate(a, b int) int {
	return a + b // Simple addition for demonstration
}