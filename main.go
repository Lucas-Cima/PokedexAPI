package main

import (
	"fmt"

	//	"centene/pokedex/mongo"
	"centene/pokedex/routes"
)

//MAIN FUNCTION
func main() {
	fmt.Println("SERVER UP")
	//	mongo.ConnectMongo()
	routes.HandleRequests()
}
