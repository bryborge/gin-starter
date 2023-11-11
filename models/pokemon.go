package models

type pokemon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TODO: Remove mocked data
var MockData = []pokemon{
	{ID: "1", Name: "Bulbasaur"},
	{ID: "2", Name: "Ivysaur"},
	{ID: "3", Name: "Venusaur"},
	{ID: "4", Name: "Charmander"},
	{ID: "5", Name: "Charmeleon"},
	{ID: "6", Name: "Charizard"},
	{ID: "7", Name: "Squirtle"},
	{ID: "8", Name: "Wartortle"},
	{ID: "9", Name: "Blastoise"},
}
