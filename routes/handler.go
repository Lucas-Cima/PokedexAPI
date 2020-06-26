package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/Lucas-Cima/PokedexAPI/model"
	"github.com/gorilla/mux"
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

func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Single Pokemon")
	tmpl := template.Must(template.ParseFiles("templates/pokemon.html"))
	vars := mux.Vars(r)
	key := vars["id"]

	for _, pokemon := range pokedex {
		if pokemon.Id == key {
			if err := tmpl.Execute(w, pokemon); err != nil {
				logrus.Error(err)
			}
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