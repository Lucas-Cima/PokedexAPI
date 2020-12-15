package model

type Pokecheck struct {
	Id     string `bson:"_id"`
	DexNum string `json:"DexNum"`
	Region string `json:"Region"`
	Name   string `json:"Name"`
	Form   string `json:"Form,omitempty"`
	Caught bool
}
