package main

import (
	"fmt"
)

func main() {
	for i := 70; i <= 100; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
