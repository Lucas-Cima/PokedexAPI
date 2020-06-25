package routes

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	//"math/rand"
	//"time"

	"github.com/gorilla/mux"

	"centene/pokedex/model"
)
//VARIABLES

var (
	newPokedex = model.PokedexService{}
	pokedex = newPokedex.GetPokemon()
)

//LANDING PAGE HANDLER
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	greeting := "Welcome to the World of Pokemon!"
	tmpl.Execute(w, greeting)
}
//POKEDEX HANDLER
func returnFullPokedex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	tmpl := template.Must(template.ParseFiles("templates/pokedex.html"))
	
	tmpl.Execute(w, pokedex)
}

//SINGLE POKEMON HANDLER
func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Single Pokemon")
	tmpl := template.Must(template.ParseFiles("templates/pokemon.html"))
	vars := mux.Vars(r)
	key := vars["id"]

	for _, pokemon := range pokedex {
		if pokemon.Id == key {
			tmpl.Execute(w, pokemon)
		}
	}
}
/*
func returnRandomPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Random Pokemon")
	tmpl := template.Must(template.ParseFiles("templates/pokemon.html"))
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(pokedex))
	vars := mux.Vars(r)
    key :=  vars["id"]

	for _, randpoke := range pokedex {
		randpoke = pokedex[randomIndex]
		if randpoke.Id == key {
			tmpl.Execute(w, randpoke)
		}
	}
}
*/	

//URL BLOCK
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokedex", returnFullPokedex).Methods("GET")
	myRouter.HandleFunc("/pokemon/{id}", returnSinglePokemon)
	//myRouter.HandleFunc("/randpoke", returnRandomPokemon)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}