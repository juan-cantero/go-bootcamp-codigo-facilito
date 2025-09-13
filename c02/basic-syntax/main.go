package main

import "fmt"

func main() {
	fmt.Println("=== GO BASIC SYNTAX EXAMPLES ===")

	// LOOPS EXAMPLES
	fmt.Println("🔄 LOOPS EXAMPLES")
	TraditionalForLoop()
	ForAsWhileLoop()
	InfiniteLoopWithBreak()
	ForRangeWithChannels()
	NestedForLoops()
	ForWithContinueAndBreak()
	ForWithCustomStep()
	ReverseForLoop()

	// CONDITIONALS EXAMPLES
	fmt.Println("🔀 CONDITIONALS EXAMPLES")
	IfStatementExamples()
	SwitchStatementExamples()

	// DATA STRUCTURES EXAMPLES
	fmt.Println("📊 DATA STRUCTURES EXAMPLES")
	ForRangeWithArrays()
	ForRangeWithSlices()
	ForRangeWithMaps()
	ForRangeWithStrings()
	ForRangeWith2DSlices()
	ForRangeWithStructSlices()
	ModifyingSliceElements()

	// FUNCTIONS EXAMPLES
	fmt.Println("⚙️ FUNCTIONS EXAMPLES")
	BasicFunction()
	demonstratePassByValue()
	demonstratePointers()
	demonstrateRecursion()

	fmt.Println("✅ All Go syntax examples completed!")
}
