package routes

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Lucas-Cima/PokedexAPI/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	findOptions := options.Find()
	findOptions.SetLimit(151)

	clientOptions := options.Client().ApplyURI("mongodb+srv://Lucas:pokemon@pokedex.l4iml.mongodb.net/Pokedex?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Pokedex").Collection("Pokemon")
	var allPokemon []model.Pokemon
	res, err := collection.Find(context.TODO(), bson.D{}, findOptions)
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

		allPokemon = append(allPokemon, pokemon)
	}
	defer func() {
		if err := res.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	tmpl := template.Must(template.ParseFiles("templates/pokedex.html"))
	if err := tmpl.Execute(w, allPokemon); err != nil {
		logrus.Error(err)
	}
}

//SINGLE POKEMON
func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Single Pokemon")
	findOptions := options.Find()
	findOptions.SetLimit(151)

	clientOptions := options.Client().ApplyURI("mongodb+srv://Lucas:pokemon@pokedex.l4iml.mongodb.net/Pokedex?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Pokedex").Collection("Pokemon")
	var allPokemon []model.Pokemon
	res, err := collection.Find(context.TODO(), bson.D{}, findOptions)
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

		allPokemon = append(allPokemon, pokemon)
	}
	defer func() {
		if err := res.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	tmpl := template.Must(template.ParseFiles("templates/pokemon.html"))
	vars := mux.Vars(r)
	key := vars["id"]

	for _, pokemon := range allPokemon {
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
	findOptions := options.Find()
	findOptions.SetLimit(151)

	clientOptions := options.Client().ApplyURI("mongodb+srv://Lucas:pokemon@pokedex.l4iml.mongodb.net/Pokedex?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Pokedex").Collection("Pokemon")
	var allPokemon []model.Pokemon
	res, err := collection.Find(context.TODO(), bson.D{}, findOptions)
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

		allPokemon = append(allPokemon, pokemon)
	}
	defer func() {
		if err := res.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	tmpl := template.Must(template.ParseFiles("templates/pokemon.html"))
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(allPokemon) - 1)
	pokemon := allPokemon[randomIndex]
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

//Who's that Pokemon!?
func whoIsDat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Who Dat!?")
	findOptions := options.Find()
	findOptions.SetLimit(151)

	clientOptions := options.Client().ApplyURI("mongodb+srv://Lucas:pokemon@pokedex.l4iml.mongodb.net/Pokedex?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Pokedex").Collection("Pokemon")
	var allPokemon []model.Pokemon
	res, err := collection.Find(context.TODO(), bson.D{}, findOptions)
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

		allPokemon = append(allPokemon, pokemon)
	}
	defer func() {
		if err := res.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	tmpl := template.Must(template.ParseFiles("templates/whodat.html"))
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(allPokemon) - 1)
	pokemon := allPokemon[randomIndex]
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
