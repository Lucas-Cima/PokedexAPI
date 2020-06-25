package main

import (
	"fmt"

	"centene/pokedex/routes"
)
//MAIN FUNCTION
func main() {
	fmt.Println("SERVER UP")
	routes.HandleRequests()
}

