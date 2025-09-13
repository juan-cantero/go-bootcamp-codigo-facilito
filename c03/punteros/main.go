package main

import "fmt"

func main() {
	x := 20
	y := 10

	fmt.Println("antes del swap")
	fmt.Println("x =", x, "y=", y)
	swap2(&x, &y)
	fmt.Println("despues del swap")
	fmt.Println("x =", x, "y=", y)
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}
