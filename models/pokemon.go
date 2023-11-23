package models

import (
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Pokemon struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func CreatePokemonInDb(driver neo4j.Driver, name string) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := session.Run(
			"CREATE (p:Pokemon {name: $name}) RETURN id(p)",
			map[string]interface{}{
				"id":   1,
				"name": name,
			},
		)

		if err != nil {
			return nil, err
		}

		return result.Consume()
	})

	return err
}

func GetPokemonInDbByName(driver neo4j.Driver, name string) (Pokemon, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		records, err := transaction.Run(
			"MATCH (p:Pokemon) WHERE p.name = $name RETURN p.name",
			map[string]interface{}{"name": name},
		)
		if err != nil {
			return nil, err
		}

		if records.Next() {
			record := records.Record()
			pokemon := Pokemon{
				Name: record.Values[0].(string),
			}
			return pokemon, nil
		}

		return nil, records.Err()
	})

	if err != nil {
		return Pokemon{}, err
	}

	if result == nil {
		return Pokemon{}, errors.New("no Pokemon found")
	}

	return result.(Pokemon), nil
}
