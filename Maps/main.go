package main

import "fmt"

func main() {
	// var colors map[string]string;
	colors := make(map[string]string);

	colors["white"] = "#ffffff";
	fmt.Println(colors);

	delete(colors, "white");
	fmt.Println(colors);

	colors = map[string]string{
		"red": "#ff0000",
		"blue": "#0000ff",
		"white": "#ffffff",
	}
	fmt.Println(colors);

	printMap(colors);
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for color", color, "is", hex);
	}
}