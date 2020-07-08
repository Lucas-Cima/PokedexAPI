package routes

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Lucas-Cima/PokedexAPI/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//VARIABLES

var (
	newPokedex = model.PokedexService{}
	pokedex    = newPokedex.GetPokemon()
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
	tmpl := template.Must(template.ParseFiles("templates/pokedex.html", "static/stylesheet.css"))
	if err := tmpl.Execute(w, pokedex); err != nil {
		logrus.Error(err)
	}
}

//SINGLE POKEMON
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

//RANDOM POKEMON
func returnRandomPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Random Pokemon")
	tmpl := template.Must(template.ParseFiles("templates/pokemon.html"))
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(pokedex) - 1)
	pokemon := pokedex[randomIndex]
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

//Who's that Pokemon!?
func whoIsDat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Who Dat!?")
	tmpl := template.Must(template.ParseFiles("templates/whodat.html"))
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(pokedex) - 1)
	pokemon := pokedex[randomIndex]
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

//Handle Requests..
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokedex", returnFullPokedex).Methods("GET")
	myRouter.HandleFunc("/pokemon/{id}", returnSinglePokemon)
	myRouter.HandleFunc("/randpoke", returnRandomPokemon)
	myRouter.HandleFunc("/whodat", whoIsDat)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}
