package model

type Pokedex interface {
	GetPokemon() (pokedex []Pokemon)
}

type PokedexService struct{}

func (p *PokedexService) GetPokemon() (pokedex []Pokemon) {
	pokedex = []Pokemon{
		{Id: "001", Name: "Bulbasaur", Element: "Grass", SecElement: "Poison", Height: "0.7 m", Weight: "6.9 kg", Species: "Seed Pokémon", Region: "Kanto", PkdxEntry: "A strange seed was planted on its back at birth. The plant sprouts and grows with this Pokémon."},
		{Id: "002", Name: "Ivysaur", Element: "Grass", SecElement: "Poison", Height: "1.0 m", Weight: "13.0 kg", Species: "Seed Pokémon", Region: "Kanto", PkdxEntry: "When the bulb on its back grows large, it appears to lose the ability to stand on its hind legs."},
		{Id: "003", Name: "Venusaur", Element: "Grass", SecElement: "Poison", Height: "2.0 m", Weight: "100.0 kg", Species: "Seed Pokémon", Region: "Kanto", PkdxEntry: "The plant blooms when it is absorbing solar energy. It stays on the move to seek sunlight."},
		{Id: "004", Name: "Charmander", Element: "Fire", SecElement: "", Height: "0.6 m", Weight: "8.5 kg", Species: "Lizard Pokémon", Region: "Kanto", PkdxEntry: "Obviously prefers hot places. When it rains, steam is said to spout from the tip of its tail."},
		{Id: "005", Name: "Charmeleon", Element: "Fire", SecElement: "", Height: "1.1 m", Weight: "19.0 kg", Species: "Flame Pokémon", Region: "Kanto", PkdxEntry: "When it swings its burning tail, it elevates the temperature to unbearably high levels."},
		{Id: "006", Name: "Charizard", Element: "Fire", SecElement: "Flying", Height: "1.7 m", Weight: "90.5 kg", Species: "Flame Pokémon", Region: "Kanto", PkdxEntry: "Spits fire that is hot enough to melt boulders. Known to cause forest fires unintentionally."},
		{Id: "007", Name: "Squirtle", Element: "Water", SecElement: "", Height: "0.5 m", Weight: "9.0 kg", Species: "Tiny Turtle Pokémon", Region: "Kanto", PkdxEntry: "After birth, its back swells and hardens into a shell. Powerfully sprays foam from its mouth."},
		{Id: "008", Name: "Wartortle", Element: "Water", SecElement: "", Height: "1.0 m", Weight: "22.5 kg", Species: "Turtle Pokémon", Region: "Kanto", PkdxEntry: "Often hides in water to stalk unwary prey. For swimming fast, it moves its ears to maintain balance."},
		{Id: "009", Name: "Blastoise", Element: "Water", SecElement: "", Height: "1.6 m", Weight: "85.5 kg", Species: "Shellfish Pokémon", Region: "Kanto", PkdxEntry: "A brutal POKéMON with pressurized water jets on its shell. They are used for high speed tackles."},
		{Id: "010", Name: "Caterpie", Element: "Bug", SecElement: "", Height: "0.3 m", Weight: "2.9 kg", Species: "Worm Pokémon", Region: "Kanto", PkdxEntry: "Its short feet are tipped with suction pads that enable it to tirelessly climb slopes and walls."},
		{Id: "011", Name: "Metapod", Element: "Bug", SecElement: "", Height: "0.7 m", Weight: "9.9 kg", Species: "Cocoon Pokémon", Region: "Kanto", PkdxEntry: "This POKéMON is vulnerable to attack while its shell is soft, exposing its weak and tender body."},
		{Id: "012", Name: "Butterfree", Element: "Bug", SecElement: "Flying", Height: "1.1 m", Weight: "32.0 kg", Species: "Butterfly Pokémon", Region: "Kanto", PkdxEntry: "In battle, it flaps its wings at high speed to release highly toxic dust into the air."},
		{Id: "013", Name: "Weedle", Element: "Bug", SecElement: "", Height: "0.3 m", Weight: "3.2 kg", Species: "Hairy Bug Pokémon", Region: "Kanto", PkdxEntry: "Often found in forests, eating leaves. It has a sharp venomous stinger on its head."},
		{Id: "014", Name: "Kakuna", Element: "Bug", SecElement: "", Height: "0.6 m", Weight: "10.0 kg", Species: "Cocoon Pokémon", Region: "Kanto", PkdxEntry: "Almost incapable of moving, this POKéMON can only harden its shell to protect itself from predators."},
		{Id: "015", Name: "Beedrill", Element: "Bug", SecElement: "Poison", Height: "1.0 m", Weight: "29.5 kg", Species: "Poison Bee Pokémon", Region: "Kanto", PkdxEntry: "Flies at high speed and attacks using its large venomous stingers on its forelegs and tail."},
	}
	return pokedex
}
