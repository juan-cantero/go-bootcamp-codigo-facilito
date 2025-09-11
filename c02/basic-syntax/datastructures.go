package main

import "fmt"

// Person struct for examples
type Person struct {
	Name string
	Age  int
}

// ForRangeWithArrays demonstrates for range with arrays
func ForRangeWithArrays() {
	fmt.Println("=== ARRAYS EXAMPLES ===")
	numbers := [5]int{10, 20, 30, 40, 50}
	
	// Get both index and value
	fmt.Println("Index and value:")
	for index, value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}
	
	// Get only index (ignore value with _)
	fmt.Println("Only indices:")
	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
	}
	
	// Get only values (ignore index with _)
	fmt.Println("Only values:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
	fmt.Println()
}

// ForRangeWithSlices demonstrates for range with slices
func ForRangeWithSlices() {
	fmt.Println("=== SLICES EXAMPLES ===")
	fruits := []string{"apple", "banana", "cherry"}
	
	fmt.Println("Original slice:")
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruit)
	}
	
	// Adding elements to slice
	fruits = append(fruits, "date", "elderberry")
	fmt.Println("After append:")
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruit)
	}
	
	// Slice operations
	fmt.Printf("Length: %d, Capacity: %d\n", len(fruits), cap(fruits))
	fmt.Printf("Slice [1:3]: %v\n", fruits[1:3])
	fmt.Println()
}

// ForRangeWithMaps demonstrates for range with maps
func ForRangeWithMaps() {
	fmt.Println("=== MAPS EXAMPLES ===")
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	
	// Get both key and value
	fmt.Println("Key and value:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
	
	// Get only keys
	fmt.Println("Only keys:")
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}
	
	// Adding and deleting from map
	ages["Diana"] = 28
	fmt.Printf("Added Diana: %v\n", ages)
	
	delete(ages, "Bob")
	fmt.Printf("Deleted Bob: %v\n", ages)
	
	// Check if key exists
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice exists and is %d years old\n", age)
	}
	fmt.Println()
}

// ForRangeWithStrings demonstrates for range with strings (runes)
func ForRangeWithStrings() {
	fmt.Println("=== STRINGS (RUNES) EXAMPLES ===")
	text := "Hello 🌎"
	
	// Get both position and rune
	fmt.Println("Position and character:")
	for pos, char := range text {
		fmt.Printf("Position %d: %c (Unicode: %d)\n", pos, char, char)
	}
	
	// String operations
	fmt.Printf("Length in bytes: %d\n", len(text))
	fmt.Printf("Length in runes: %d\n", len([]rune(text)))
	fmt.Println()
}

// ForRangeWith2DSlices demonstrates 2D slices
func ForRangeWith2DSlices() {
	fmt.Println("=== 2D SLICES EXAMPLES ===")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	fmt.Println("Matrix:")
	for rowIndex, row := range matrix {
		fmt.Printf("Row %d: ", rowIndex)
		for colIndex, value := range row {
			fmt.Printf("[%d,%d]=%d ", rowIndex, colIndex, value)
		}
		fmt.Println()
	}
	fmt.Println()
}

// ForRangeWithStructSlices demonstrates struct slices
func ForRangeWithStructSlices() {
	fmt.Println("=== STRUCT SLICES EXAMPLES ===")
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	
	fmt.Println("People:")
	for index, person := range people {
		fmt.Printf("Person %d: %s (%d years old)\n", index+1, person.Name, person.Age)
	}
	
	// Adding a new person
	people = append(people, Person{Name: "Diana", Age: 28})
	fmt.Printf("Added Diana: %+v\n", people[len(people)-1])
	fmt.Println()
}

// ModifyingSliceElements demonstrates the difference between modifying by value vs index
func ModifyingSliceElements() {
	fmt.Println("=== MODIFYING SLICE ELEMENTS ===")
	nums := []int{1, 2, 3, 4, 5}
	
	// This DOES NOT modify the original slice (value is a copy)
	fmt.Println("Before (incorrect way):")
	fmt.Printf("Original: %v\n", nums)
	for _, value := range nums {
		value *= 2 // This doesn't change the original
	}
	fmt.Printf("After range with value: %v\n", nums) // Still [1, 2, 3, 4, 5]
	
	// This DOES modify the original slice (using index)
	fmt.Println("After (correct way):")
	for i := range nums {
		nums[i] *= 2 // This changes the original
	}
	fmt.Printf("After range with index: %v\n", nums) // Now [2, 4, 6, 8, 10]
	fmt.Println()
}