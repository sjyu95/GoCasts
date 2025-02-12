package main

import "fmt"

func main() {
	var int_arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, e := range int_arr {
		if e%2 == 0 {
			fmt.Println(e, " is even")
		} else {
			fmt.Println(e, " is odd")
		}
	}
}
