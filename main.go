package main

import (
	"fmt"

	"github.com/Lucas-Cima/PokedexAPI/routes"
)
//MAIN FUNCTION
func main() {
	fmt.Println("SERVER UP")
	routes.HandleRequests()
}

