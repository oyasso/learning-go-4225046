package main

import (
	"fmt"
)

func main() {
	doSomething()
	value1 := 5
	value2 := 10
	value3 := 65
	sum, count, average := addAllValues(value1, value2, value3)
	fmt.Printf("The sum is %v\n", sum)
	fmt.Printf("The count is %v\n", count)
	fmt.Printf("The average is %v\n", average)

}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int, float64) {
	sum := 0
	for _, v := range values {
		sum += v
	}			
	count := len(values)
	average := float64(sum) / float64(count)	
	return sum, count, average
}