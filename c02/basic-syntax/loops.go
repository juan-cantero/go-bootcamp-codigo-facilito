package main

import "fmt"

// TraditionalForLoop demonstrates C-style for loops
func TraditionalForLoop() {
	fmt.Println("1. Traditional for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	fmt.Println()
}

// ForAsWhileLoop demonstrates for loops used as while loops
func ForAsWhileLoop() {
	fmt.Println("2. For as while loop:")
	counter := 0
	for counter < 3 { // No init or increment - acts like while
		fmt.Printf("counter = %d\n", counter)
		counter++
	}
	fmt.Println()
}

// InfiniteLoopWithBreak demonstrates infinite loops with break
func InfiniteLoopWithBreak() {
	fmt.Println("3. Infinite loop with break:")
	n := 0
	for n < 3 { // Modified to be a while-style loop
		fmt.Printf("n = %d\n", n)
		n++
	}
	fmt.Println()
}

// ForRangeWithChannels demonstrates for range with channels
func ForRangeWithChannels() {
	fmt.Println("4. For range with channels:")
	ch := make(chan int, 3) // Buffered channel
	
	// Send values to channel
	ch <- 100
	ch <- 200
	ch <- 300
	close(ch) // Must close to end the range loop
	
	// Range over channel - receives values until closed
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println()
}

// NestedForLoops demonstrates nested for loops
func NestedForLoops() {
	fmt.Println("5. Nested for loops (multiplication table):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println() // New line after each row
	}
	fmt.Println()
}

// ForWithContinueAndBreak demonstrates continue and break
func ForWithContinueAndBreak() {
	fmt.Println("6. For with continue and break:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		if i > 7 {
			break // Stop when i > 7
		}
		fmt.Printf("Odd number: %d\n", i)
	}
	fmt.Println()
}

// ForWithCustomStep demonstrates custom increments
func ForWithCustomStep() {
	fmt.Println("7. For with custom step:")
	for i := 0; i < 20; i += 3 { // Increment by 3
		fmt.Printf("i = %d\n", i)
	}
	fmt.Println()
}

// ReverseForLoop demonstrates reverse iteration
func ReverseForLoop() {
	fmt.Println("8. Reverse for loop:")
	for i := 10; i >= 0; i-- { // Decrement
		fmt.Printf("Countdown: %d\n", i)
	}
	fmt.Println()
}