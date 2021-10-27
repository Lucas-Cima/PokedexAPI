package routes

import (
	"context"
	"encoding/json"
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
	MongoDb3 mongo.Collection
	header   = "templates/header.html"
	pokecard = "templates/pokecard.html"
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

//GET FULL LIST OF TRAINERS
func getTrainerList(collection *mongo.Collection) (trainers []model.Trainer) {
	res, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logrus.Error(err)
	}
	for res.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var trainer model.Trainer
		err := res.Decode(&trainer)
		if err != nil {
			log.Fatal(err)
		}
		trainers = append(trainers, trainer)
	}
	return
}

//LANDING PAGE HANDLER
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	tmpl := template.Must(template.ParseFiles("templates/index.html", header))
	greeting := "Welcome to the World of Pokemon!"
	if err := tmpl.Execute(w, greeting); err != nil {
		logrus.Error(err)
	}
}

//POKEDEX HANDLER
func returnFullPokedex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endoint Hit: Full Pokedex")
	tmpl := template.Must(template.ParseFiles("templates/pokedex/pokedex.html", header, pokecard))
	pokedex := getPokedex(&MongoDb)
	if err := tmpl.Execute(w, pokedex); err != nil {
		logrus.Error(err)
	}
}

//SINGLE POKEMON
func returnSinglePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Single Pokemon")
	tmpl := template.Must(template.ParseFiles("templates/pokedex/pokemon.html", header, pokecard))
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
	tmpl := template.Must(template.ParseFiles("templates/pokedex/pokemon.html", header, pokecard))
	pokemon := getRandom(&MongoDb)
	if err := tmpl.Execute(w, pokemon); err != nil {
		logrus.Error(err)
	}
}

//Who's that Pokemon!?
func whoIsDat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Endpoint Hit: Who Dat!?")
		tmpl := template.Must(template.ParseFiles("templates/whodat/who-dat.html", header))
		greeting := "WHO's THAT POKEMON!?"
		if err := tmpl.Execute(w, greeting); err != nil {
			logrus.Error(err)
		}
	}
	if r.Method == "POST" {
		//EASY MODE
		if r.URL.String() == "/whodatEasy" {
			fmt.Println("Endpoint Hit: Who Dat!?...Easy Mode")
			tmpl := template.Must(template.ParseFiles("templates/whodat/whodatEasy.html", header))
			pokemon := getRandom(&MongoDb)
			fmt.Println(pokemon.Name + " " + pokemon.Id)
			if err := tmpl.Execute(w, pokemon); err != nil {
				logrus.Error(err)
			}
		}
		//MEDIUM MODE
		if r.URL.String() == "/whodatMedium" {
			fmt.Println("Endpoint Hit: Who Dat!?...Medium Mode")
			tmpl := template.Must(template.ParseFiles("templates/whodat/whodatMedium.html", header))
			pokemon := getRandom(&MongoDb)
			fmt.Println(pokemon.Name + " " + pokemon.Id)
			if err := tmpl.Execute(w, pokemon); err != nil {
				logrus.Error(err)
			}
		}
		//HARD MODE
		if r.URL.String() == "/whodatHard" {
			fmt.Println("Endpoint Hit: Who Dat!?...Hard Mode")
			tmpl := template.Must(template.ParseFiles("templates/whodat/whodatHard.html", header))
			pokemon := getRandom(&MongoDb)
			fmt.Println(pokemon.Name + " " + pokemon.Id)
			if err := tmpl.Execute(w, pokemon); err != nil {
				logrus.Error(err)
			}
		}
	}
}

//RETURNS LIST OF TRAINERS
func returnTrainerList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Trainer List")
	tmpl := template.Must(template.ParseFiles("templates/trainer/trainerlist.html", header))
	trainers := getTrainerList(&MongoDb3)
	if err := tmpl.Execute(w, trainers); err != nil {
		logrus.Error(err)
	}
}

//RETURNS SINGLE TRAINER
func returnSingleTrainer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Single Trainer")
	tmpl := template.Must(template.ParseFiles("templates/trainer/trainer.html", header))
	trainerList := getTrainerList(&MongoDb3)
	vars := mux.Vars(r)
	key := vars["name"]
	for _, trainer := range trainerList {
		if trainer.Name == key {
			if err := tmpl.Execute(w, trainer); err != nil {
				logrus.Error(err)
			}
		}
	}
}

func returnCreateTrainer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Endpoint Hit: Create Trainer")
		tmpl := template.Must(template.ParseFiles("templates/trainer/createtrainer.html", header))
		if err := tmpl.Execute(w, nil); err != nil {
			logrus.Error(err)
		}
	}
	if r.Method == "POST" {
		fmt.Println("Endpoint Hit: Trainer Created")
		var trainer model.Trainer
		err := json.NewDecoder(r.Body).Decode(&trainer)
		logrus.Info(trainer)
		if err != nil {
			logrus.Error(err)
		}
		insertResult, err := MongoDb3.InsertOne(r.Context(), trainer)
		if err != nil {
			logrus.Error(err)
		}
		json.NewEncoder(w).Encode(insertResult.InsertedID)
		fmt.Println(r.Body)
	}
}

//Handle Requests
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/pokedex", returnFullPokedex)
	myRouter.HandleFunc("/pokemon/{id}", returnSinglePokemon)
	myRouter.HandleFunc("/randpoke", returnRandomPokemon)
	myRouter.HandleFunc("/whodat", whoIsDat)
	myRouter.HandleFunc("/whodatEasy", whoIsDat)
	myRouter.HandleFunc("/whodatMedium", whoIsDat)
	myRouter.HandleFunc("/whodatHard", whoIsDat)
	myRouter.HandleFunc("/trainerList", returnTrainerList)
	myRouter.HandleFunc("/trainer/{name}", returnSingleTrainer)
	myRouter.HandleFunc("/createTrainer", returnCreateTrainer)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}
