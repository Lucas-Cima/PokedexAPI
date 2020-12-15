package routes

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Lucas-Cima/PokedexAPI/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoDb  mongo.Collection
	MongoDb2 mongo.Collection
)

//Getting a random pokemon
func getRandom(collection *mongo.Collection) (randPoke model.Pokemon) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(151)
	i := strconv.Itoa(randomIndex)
	if randomIndex < 100 && randomIndex >= 10 {
		i = fmt.Sprintf("0%v", i)
	} else if randomIndex < 10 {
		i = fmt.Sprintf("00%v", i)
	}
	if err := collection.FindOne(context.Background(), bson.M{"_id": i}).Decode(&randPoke); err != nil {
		logrus.Errorf("%v: index: %v", err, i)
	}
	return
}

//Getting Full Pokedex
func getPokedex(collection *mongo.Collection) (pokedex []model.Pokemon) {
	res, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logrus.Error(err)
	}
	for res.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var pokemon model.Pokemon
		err := res.Decode(&pokemon)
		if err != nil {
			log.Fatal(err)
		}
		pokedex = append(pokedex, pokemon)
	}
	return
}

func getPokeCheck(collection *mongo.Collection) (pokedex []model.Pokecheck) {
	res, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logrus.Error(err)
	}
	for res.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var pokecheck model.Pokecheck
		err := res.Decode(&pokecheck)
		if err != nil {
			log.Fatal(err)
		}
		pokedex = append(pokedex, pokecheck)
	}
	return
}

//LANDING PAGE HANDLER
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	tmpl := template.Must(template.ParseFiles("pokedexAPI/src/app/app.component.html"))
	greeting := "Welcome to the World of Pokemon!"
	if err := tmpl.Execute(w, greeting); err != nil {
		logrus.Error(err)
	}
}

//POKEDEX HANDLER
func returnFullPokedex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	tmpl := template.Must(template.ParseFiles("pokedexAPI/src/app/pokedex/pokedex.component.html"))
	pokedex := getPokedex(&MongoDb)
	if err := tmpl.Execute(w, pokedex); err != nil {
		logrus.Error(err)
	}
}

//SINGLE POKEMON
func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Single Pokemon")
	tmpl := template.Must(template.ParseFiles("pokedexAPI/src/app/pokemon/pokemon.component.html"))
	pokedex := getPokedex(&MongoDb)
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
	tmpl := template.Must(template.ParseFiles("pokedexAPI/src/app/pokemon/pokemon.component.html"))
	pokemon := getRandom(&MongoDb)
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

//Who's that Pokemon!?
func whoIsDat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Who Dat!?")
	tmpl := template.Must(template.ParseFiles("pokedexAPI/src/app/who-dat/who-dat.component.html"))
	greeting := "WHO's THAT POKEMON!?"
	if err := tmpl.Execute(w, greeting); err != nil {
		logrus.Error(err)
	}
}

func whoIsDatEasy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Who Dat!?...Easy Mode")
	tmpl := template.Must(template.ParseFiles("templates/whodatEasy.html"))
	pokemon := getRandom(&MongoDb)
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

func whoIsDatMedium(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Who Dat!?...Medium Mode")
	tmpl := template.Must(template.ParseFiles("templates/whodatMedium.html"))
	pokemon := getRandom(&MongoDb)
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

func whoIsDatHard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Who Dat!?...Hard Mode")
	tmpl := template.Must(template.ParseFiles("templates/whodatHard.html"))
	pokemon := getRandom(&MongoDb)
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

func pokemonCheckList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Checklist")
	tmpl := template.Must(template.ParseFiles("pokedexAPI/src/app/pokecheck/pokecheck.component.html"))
	pokecheck := getPokeCheck(&MongoDb2)
	if err := tmpl.Execute(w, pokecheck); err != nil {
		logrus.Error(err)
	}
}

//Handle Requests..
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokedex", returnFullPokedex)
	myRouter.HandleFunc("/pokemon/{id}", returnSinglePokemon)
	myRouter.HandleFunc("/randpoke", returnRandomPokemon)
	myRouter.HandleFunc("/whodat", whoIsDat)
	myRouter.HandleFunc("/whodatEasy", whoIsDatEasy)
	myRouter.HandleFunc("/whodatMedium", whoIsDatMedium)
	myRouter.HandleFunc("/whodatHard", whoIsDatHard)
	myRouter.HandleFunc("/pokecheck", pokemonCheckList)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}
