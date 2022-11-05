package core

import (
	"encoding/json"
	"os"
	"sync"
)

var fileLock = &sync.RWMutex{}

type CounterData struct {
	Pid     int
	Counter int
}

func LoadCounterData(path string) (*CounterData, error) {
	fileLock.RLock()
	defer fileLock.RUnlock()

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
	fileLock.Lock()
	defer fileLock.Unlock()

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	err = json.NewEncoder(file).Encode(counterData)
	return err
}

func EditCounterData(path string, action func(*CounterData)) (*CounterData, error) {
	fileLock.Lock()
	defer fileLock.Unlock()

	file, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	counterData := &CounterData{}
	if err = json.NewDecoder(file).Decode(counterData); err != nil {
		return nil, err
	}
	action(counterData)
	err = json.NewEncoder(file).Encode(counterData)

	return counterData, err
}
