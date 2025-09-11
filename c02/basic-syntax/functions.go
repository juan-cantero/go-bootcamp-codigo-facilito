package main

import (
	"fmt"
	"errors"
)

// BasicFunction demonstrates a simple function
func BasicFunction() {
	fmt.Println("=== FUNCTIONS EXAMPLES ===")
	
	// Call functions with different patterns
	result := add(5, 3)
	fmt.Printf("add(5, 3) = %d\n", result)
	
	// Multiple return values
	sum, diff, prod, quot := mathOperations(10, 5)
	fmt.Printf("mathOperations(10, 5) = sum:%d, diff:%d, prod:%d, quot:%.2f\n", sum, diff, prod, quot)
	
	// Named return values
	q, r := divideWithRemainder(17, 5)
	fmt.Printf("divideWithRemainder(17, 5) = quotient:%d, remainder:%d\n", q, r)
	
	// Error handling
	result2, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("safeDivide(10, 2) = %.2f\n", result2)
	}
	
	result3, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("safeDivide(10, 0) = %.2f\n", result3)
	}
	
	// Variadic functions
	total1 := sumNumbers(1, 2, 3, 4, 5)
	fmt.Printf("sumNumbers(1, 2, 3, 4, 5) = %d\n", total1)
	
	numbers := []int{10, 20, 30}
	total2 := sumNumbers(numbers...) // Spread slice
	fmt.Printf("sumNumbers(10, 20, 30) = %d\n", total2)
	
	// Anonymous functions and closures
	demonstrateClosures()
	
	fmt.Println()
}

// Simple function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func mathOperations(a, b int) (int, int, int, float64) {
	sum := a + b
	difference := a - b
	product := a * b
	quotient := float64(a) / float64(b)
	return sum, difference, product, quotient
}

// Function with named return values
func divideWithRemainder(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return // Named return - automatically returns quotient and remainder
}

// Function that returns an error
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Variadic function (accepts variable number of arguments)
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function that demonstrates closures
func demonstrateClosures() {
	fmt.Println("--- Closures Examples ---")
	
	// Function that returns a function (closure)
	counter := makeCounter()
	fmt.Printf("Counter 1: %d\n", counter()) // 1
	fmt.Printf("Counter 1: %d\n", counter()) // 2
	fmt.Printf("Counter 1: %d\n", counter()) // 3
	
	// Another counter instance
	counter2 := makeCounter()
	fmt.Printf("Counter 2: %d\n", counter2()) // 1 (independent)
	
	// Anonymous function
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Anonymous square(4) = %d\n", square(4))
	
	// Immediately invoked function
	result := func(a, b int) int {
		return a * b
	}(6, 7)
	fmt.Printf("IIFE result = %d\n", result)
}

// Function that returns a function (closure)
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Function with different parameter types
func processUser(name string, age int, isActive bool) string {
	status := "inactive"
	if isActive {
		status = "active"
	}
	return fmt.Sprintf("User %s (age %d) is %s", name, age, status)
}

// Function that modifies a slice (reference type)
func modifySlice(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

// Function that demonstrates pass by value vs reference
func demonstratePassByValue() {
	fmt.Println("--- Pass by Value vs Reference ---")
	
	// Pass by value (primitive types)
	x := 10
	fmt.Printf("Before doubleValue: x = %d\n", x)
	doubleValue(x)
	fmt.Printf("After doubleValue: x = %d\n", x) // Still 10
	
	// Pass by reference (slices, maps, channels)
	nums := []int{1, 2, 3}
	fmt.Printf("Before modifySlice: nums = %v\n", nums)
	modifySlice(nums)
	fmt.Printf("After modifySlice: nums = %v\n", nums) // Modified
}

// Function that takes value (doesn't modify original)
func doubleValue(n int) {
	n = n * 2
	fmt.Printf("Inside doubleValue: n = %d\n", n)
}

// Function that takes a pointer (can modify original)
func doublePointer(n *int) {
	*n = *n * 2
}

// Function demonstrating pointers
func demonstratePointers() {
	fmt.Println("--- Pointers Examples ---")
	
	x := 10
	fmt.Printf("Before doublePointer: x = %d\n", x)
	doublePointer(&x) // Pass address of x
	fmt.Printf("After doublePointer: x = %d\n", x) // Now 20
}

// Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Function demonstrating recursion
func demonstrateRecursion() {
	fmt.Println("--- Recursion Examples ---")
	
	for i := 0; i <= 5; i++ {
		result := factorial(i)
		fmt.Printf("factorial(%d) = %d\n", i, result)
	}
}