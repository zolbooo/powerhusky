package core

import (
	"encoding/json"
	"os"
)

type CounterData struct {
	Pid     int
	Counter int
}

func LoadCounterData(path string) (*CounterData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result := &CounterData{}
	if err = json.NewDecoder(file).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (counterData *CounterData) Save(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	err = json.NewEncoder(file).Encode(counterData)
	return err
}
