package main

import (
	"fmt"

	"github.com/Lucas-Cima/PokedexAPI/routes"

	//	"centene/pokedex/mongo"
)

//MAIN FUNCTION
func main() {
	fmt.Println("SERVER UP")
	//	mongo.ConnectMongo()
	routes.HandleRequests()
}
