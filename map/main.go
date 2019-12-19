package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4BF745",
		"white": "#FFFFFF",
	}

	/* Have to use bracket notation with maps. Cannot use dot notation. */
	// colors[10] = "#FFFFFF"

	/* How to delete a map */
	// delete(colors, 10)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

/* Two other ways of making a map */
// var colors map[string]string
// colors := make(map[int]string)
