package main

import "fmt"

func main() {
	var First, Second bool
	var Third bool = true
	var Fours = !Third
	var Fifth = true

	fmt.Println("First  = ", First)
	fmt.Println("Second = ", Second)
	fmt.Println("Third  = ", Third)
	fmt.Println("Fours  = ", Fours)
	fmt.Println("Fifth  = ", Fifth, "\n")

	fmt.Println("!true = ", !true)
	fmt.Println("!false = ", !false, "\n")

	fmt.Println("true && true = ", true && true)
	fmt.Println("true && false = ", true && false)
	fmt.Println("false && false =", false && false, "\n")

	fmt.Println("true || true = ", true || true)
	fmt.Println("true || false = ", true || false)
	fmt.Println("false || false =", false || false, "\n")
	
	fmt.Println("2 < 3 = ", 2<3)
	fmt.Println("2 > 3 = ", 2>3)
	fmt.Println("3 < 3 = ", 3<3)
	fmt.Println("3 <= 3 = ", 3<=3)
	fmt.Println("3 > 3 = ", 3>3)
	fmt.Println("3 >= 3 = ", 3>=3)
	fmt.Println("2 == 3 = ", 2==3)
	fmt.Println("3 == 3 = ", 3==3)
	fmt.Println("2 != 3 = ", 2!=3)
	fmt.Println("3 != 3 = ", 3!=3)
	
}