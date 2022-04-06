package main

import (
	"fmt"
	"sort"
	"strings"
)

type user struct {
	name    string
	surname string
}
type player struct {
	name  string
	score int
}

func main() {
	// Declare and initialize the map with values.
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}
	fmt.Println("Sort the map by key:")
	// pull the keys from the map
	var keys []string
	for k := range users {
		keys = append(keys, k)
	}
	// sort the keys
	sort.Strings(keys)

	for _, k := range keys {
		// pull the values from the map
		u := users[k]
		// print the values
		fmt.Println(k, u.name, u.surname)
	}

	fmt.Printf("\nYou cannot take the address of an element in a map\n")
	fmt.Printf("You can only take the address of the map itself\n")
	players := map[string]player{
		"anna": {"Anna", 42},
		"bob":  {"Bob", 12},
	}
	//anna := &players["anna"]
	//anna.score++
	//invalid operation: cannot take address of players["anna"] (map index expression of type player)

	// instead take the element, modify it, and put it back
	player2 := players["anna"]
	player2.score++
	players["anna"] = player2
	fmt.Println(players)

	fmt.Println(strings.Repeat("-", 40))
	mapRef()
	fmt.Println(strings.Repeat("-", 40))
	mapRef2()
}

// Sample program to show how maps are reference types.
func mapRef() {
	// Declare and make a map with values.
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Make a copy of the map.
	users2 := users

	// Delete an element from the map.
	delete(users, "Mouse")

	// Iterate over the map.
	for k, v := range users {
		fmt.Println(k, v.name, v.surname)
	}

	// Iterate over the copy of the map.
	for k, v := range users2 {
		fmt.Println(k, v.name, v.surname)
	}
}

func mapRef2() {
	// Initialize a map with values.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Pass the map to a function to perform some mutation.
	double(scores, "anna")

	// See the change is visible in our map.
	fmt.Println("Score:", scores["anna"])
}

func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}
