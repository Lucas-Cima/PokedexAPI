package main

import (
	"fmt"

	//	"centene/pokedex/mongo"
	"github.com/Lucas-Cima/PokedexAPI/routes"
)

//MAIN FUNCTION
func main() {
	fmt.Println("SERVER UP")
	//	mongo.ConnectMongo()
	routes.HandleRequests()
}
