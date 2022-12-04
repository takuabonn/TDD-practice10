package main

import "fmt"

func numberOfShouts(a,b,k int) int {
	i := 1
	resultKx := k 

	if a == b {
		return 0
	}

	for a * resultKx < b {
	    resultKx = resultKx * k
		i++
	}
    return i
}

func main() {
    // fmt.Println(numberOfShouts())
}