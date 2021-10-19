package model

type PokemonList struct {
	DexNum string `json:"DexNum"`
	Name   string `json:"Name"`
	Form   string `json:"Form,omitempty"`
}
