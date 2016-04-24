package main

import (
	"fmt"
	"log"
)

// IsEven checks if the number is even or not
// This example shows how to add more info to `error`.
func IsEven(n int) error {
	if n%2 != 0 {
		return fmt.Errorf("%d is not even", n)
	}
	return nil
}

// IsDevisibleBy checks if x is devisible by y
func IsDevisibleBy(x, y int) error {
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range nums {
		if err := IsEven(n); err != nil {
			log.Println(err)
		}
	}
}
