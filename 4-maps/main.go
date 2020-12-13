package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"grenn": "4567454",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("The value for ", key, " is ", value)
	}
}
