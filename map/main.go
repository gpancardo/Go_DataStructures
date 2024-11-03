package main

import "fmt"

func main() {
	var user_color string
	var rm string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#0000ff",
		"blue":  "#00ff00",
		"white": "#000000",
		"black": "#FFFFFF",
	}
	fmt.Println("What color to consult? ")
	fmt.Scan(&user_color)
	fmt.Println("You asked if ", user_color, " is on our map")
	//exists is a variable that Go creates by default when a map element is queried. It's a boolean
	search, exists := colors[user_color]
	fmt.Println(exists, search)
	fmt.Println("Want to delete a color? ")
	fmt.Scan(&rm)
	if rm == "yes" {
		current_long := len(colors)
		if current_long > 0 {
			fmt.Printf("There are %d colors", current_long)
			delete(colors, "green")
			fmt.Printf("\nNow there are %d colors", len(colors))
		}

	}
}
