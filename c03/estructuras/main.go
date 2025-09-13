package main

import "fmt"

type NamePrintable interface {
	PrintName()
}

type Address struct {
	Num  string
	Name string
}

type User struct {
	Id      int
	Name    string
	Address Address
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

type Professor struct {
	Id   int
	Name string
}

func (p Professor) PrintName() {
	fmt.Println(p.Name)
}

func PrintStrucName(u NamePrintable) {
	u.PrintName()
}

func main() {
	user := User{Id: 1, Name: "Juan", Address: Address{Num: "123465", Name: "camargo"}}

	fmt.Println(user)

	professor := Professor{Id: 1, Name: "John"}

	PrintStrucName(professor)
}
