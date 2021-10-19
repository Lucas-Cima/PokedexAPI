package model

type Trainer struct {
	Id          string `bson:"_id"`
	Name        string `json:"Name"`
	PokemonList []PokemonList
}
