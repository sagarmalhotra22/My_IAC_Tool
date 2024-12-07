package state

import (
	"encoding/json"
	"log"
	"os"
)

var StateFile = "state.json"

type State struct {
	Instances string `json:"instances"`
	ProjectId string `json:"projectId"`
	Zone      string `json:"zone"`
}

func SaveState(state *State) error {
	file, err := os.Create(StateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(state)
}

func LoadState() (*State, error) {
	file, err := os.Open(StateFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Failed to load state: No state file")
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

// DeleteState removes the state file.
func DeleteState() error {
	if err := os.Remove(StateFile); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}
