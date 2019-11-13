package main

import "fmt"
//https://gobyexample.com/for

// Note: for is Go’s only looping construct. Here are three basic types of for loops.
func main()  {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n %2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}