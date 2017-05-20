package main

import "fmt"

func main () {

x := []int {6,4,8,1,4,6,10,}
lowest := 99 

    for i:= 0; i < len(x); i++ {

        if x[i] < lowest {
            lowest = x[i];
        }
    }

fmt.Println(lowest)
}

