package routes

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"log"
	"net/http"

	"github.com/Lucas-Cima/PokedexAPI/model"
	"github.com/gorilla/mux"
)

var (
	newPokedex = model.PokedexService{}
	pokedex = newPokedex.GetPokemon()
)

//LANDING PAGE HANDLER
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	greeting := "Welcome to the World of Pokemon!"
	if err := tmpl.Execute(w, greeting); err != nil {
		logrus.Error(err)
	}
}
//POKEDEX HANDLER
func returnFullPokedex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	tmpl := template.Must(template.ParseFiles("templates/pokedex.html"))

	if err := tmpl.Execute(w, pokedex); err != nil {
		logrus.Error(err)
	}
}

//SINGLE POKEMON HANDLER
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

//URL BLOCK
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokedex", returnFullPokedex).Methods("GET")
	myRouter.HandleFunc("/pokemon/{id}", returnSinglePokemon)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}