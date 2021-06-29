package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canMyAgeDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi, ramen"}
	nico1 := person{"nico", 18, favFood}
	// nico2 := person{name: "nico", age: 18, favFood} // mixture of field:value and value elements in struct literal
	nico3 := person{name: "nico", favFood: favFood, age: 18}
	fmt.Println(nico1, nico3.name)
}
