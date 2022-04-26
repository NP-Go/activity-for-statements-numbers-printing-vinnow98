package main

import "fmt"

func activity1(x int) int {
	for x := 0; x <= 1000; x++ {
		fmt.Println(x)
	}
	return x
}

func main() {
	var x int
	activity1(x)

}
