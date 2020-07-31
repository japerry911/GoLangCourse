package main

import "fmt"

func main() {
	var numbers [11]int

	for i := range numbers {
		if i%2 == 0 {
			fmt.Printf("%d is Even\n", i)
		} else {
			fmt.Printf("%d is Odd\n", i)
		}
	}
}
