package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[int]string)

	// colors := map[string]string {
	// 	"red": "#FF0000",
	// 	"green": "#4BF745",
	// }

	colors[10] = "#FFFFFF" //Have to use bracket notation with maps. Cannot use dot notation.

	delete(colors, 10)

	fmt.Println(colors)
}
