package match

import (
	"encoding/json"
	"os"
)

type repository struct{}

var instance *repository

func Repository() *repository {
	if instance == nil {
		instance = &repository{}
	}

	return instance
}

func (r *repository) Get(id string) (*Match, error) {
	type matchJSON struct {
		ID    string          `json:"id"`
		Title string          `json:"title"`
		Facts json.RawMessage `json:"facts"`
	}

	data, err := os.ReadFile("./json/match.json")

	if err != nil {
		return nil, err
	}

	var m matchJSON

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return New(m.ID, m.Title, m.Facts), nil
}
