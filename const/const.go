package main

import "fmt"

const ( 
	First = 1
	Second
	Third = "Third"
	Fours
)

const ( 
	IotaFirst = iota
	IotaSecond
	IotaThird
	IotaFours
)

const ( 
	IotabitFirst = 1 << iota 
	IotabitSecond
	IotabitThird
	IotabitFours
)

const ( 
	IotaCaseFirst = iota * 3 - 2
	IotaCaseSecond
	IotaCaseThird
	IotaCaseFours
)

func main()  {
	fmt.Println("First  = ", First)
	fmt.Println("Second = ", Second)
	fmt.Println("Third  = ", Third)
	fmt.Println("Fours  = ", Fours, "\n")

	fmt.Println("Iota First  = ", IotaFirst)
	fmt.Println("Iota Second = ", IotaSecond)
	fmt.Println("Iota Third  = ", IotaThird)
	fmt.Println("Iota Fours  = ", IotaFours, "\n")

	fmt.Println("Iota bit First  = ", IotabitFirst)
	fmt.Println("Iota bit Second = ", IotabitSecond)
	fmt.Println("Iota bit Third  = ", IotabitThird)
	fmt.Println("Iota bits Fours = ", IotabitFours, "\n")

	fmt.Println("Iota case First  =", IotaCaseFirst)
	fmt.Println("Iota case Second = ", IotaCaseSecond)
	fmt.Println("Iota case Third  = ", IotaCaseThird)
	fmt.Println("Iota case Fours  = ", IotaCaseFours, "\n")
}