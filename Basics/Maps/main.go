package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	fmt.Println(colors3)

	colors["white"] = "#FFFFFF"
	delete(colors, "red")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}
