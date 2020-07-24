package model

type Pokemon struct {
	Id         string `bson:"_id"`
	Name       string `json:"Name"`
	Element    string `json:"Element"`
	SecElement string `json:"SecElement,omitempty"`
	Height     string `json:"Height"`
	Weight     string `json:"Weight"`
	Species    string `json:"Species"`
	Region     string `json:"Region"`
	PkdxEntry  string `json:"PkdxEntry"`
}
