package main

import "fmt"

func main () {

x := []int {6,4,8,4,6,10,}
var min int

    for i, value := range x {
	if i==0 || value<= min {
           min = value
	}
    }

fmt.Println(min)
}

