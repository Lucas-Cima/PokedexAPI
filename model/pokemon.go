package model

type Pokemon struct {
	Id         string `bson:"Id"`
	Name       string `bson:"Name"`
	Element    string `bson:"Element"`
	SecElement string `bson:"SecElement,omitempty"`
	Height     string `bson:"Height"`
	Weight     string `bson:"Weight"`
	Species    string `bson:"Species"`
	Region     string `bson:"Region"`
	PkdxEntry  string `bson:"PkdxEntry"`
}
