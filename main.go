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

func main() {
	totalLenght, upperName := lenAndUpper("henry")
	fmt.Println(totalLenght, upperName)
	// totalLenght, _ := lenAndUpper("henry")
	// repeatMe("a", "b", "c")

}
