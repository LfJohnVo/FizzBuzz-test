package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hola mundo")
	fizzbuzz(15)

}

func fizzbuzz(n int32){
	for i := int32(1); i <= n; i++ {

		if i%3 == 0 {
			// Multiple of 3
			fmt.Printf("fizz")
		}
		if i%5 == 0 {
			// Multiple of 5
			fmt.Printf("buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			// Neither, so print the number itself
			fmt.Printf("%d", i)
		}

		// A trailing new line (so both fizz + buzz can be printed on the same line)
		fmt.Printf("\n")

	}
}