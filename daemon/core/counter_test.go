package core

import (
	"os"
	"testing"
)

func TestCounter(t *testing.T) {
	tmpfile, err := os.CreateTemp(os.TempDir(), "*.pid")
	if err != nil {
		t.Fatalf("failed to create temp file: %+v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	counterData := &CounterData{Counter: 3, Pid: os.Getpid()}
	counterData.Save(tmpfile.Name())

	newCounterData, err := LoadCounterData(tmpfile.Name())
	if err != nil {
		t.Errorf("failed to load counter data: %+v", err)
	}
	if newCounterData.Counter != 3 || newCounterData.Pid != os.Getpid() {
		t.Errorf("invalid counter data: %+v", newCounterData)
	}
}

func TestEditCounter(t *testing.T) {
	tmpfile, err := os.CreateTemp(os.TempDir(), "*.pid")
	if err != nil {
		t.Fatalf("failed to create temp file: %+v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	counterData := &CounterData{Counter: 3, Pid: os.Getpid()}
	counterData.Save(tmpfile.Name())

	updatedData, err := EditCounterData(tmpfile.Name(), func(cd *CounterData) {
		cd.Counter += 10
	})
	if err != nil {
		t.Errorf("failed to edit counter data: %+v", err)
	}

	newCounterData, err := LoadCounterData(tmpfile.Name())
	if err != nil {
		t.Errorf("failed to load counter data: %+v", err)
	}
	if newCounterData.Counter != updatedData.Counter || newCounterData.Counter != 13 || newCounterData.Pid != os.Getpid() {
		t.Errorf("invalid counter data: %+v, updated data was: %+v", newCounterData, updatedData)
	}
}
