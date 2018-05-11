package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, i, Pow(i))
	}

}
