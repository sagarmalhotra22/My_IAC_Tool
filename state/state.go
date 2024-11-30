package state

import (
	"encoding/json"
	"os"
)

type State struct {
	Instances []string `json:"instances"`
}

func SaveState(path string, state *State) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(state)
}

func LoadState(path string) (*State, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &State{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var state State
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&state); err != nil {
		return nil, err
	}

	return &state, nil
}
