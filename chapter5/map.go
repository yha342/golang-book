package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"

	name, ok := elements["He"]
	fmt.Println(name, ok)
	}
