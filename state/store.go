package state

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type State struct {
	StartedAt time.Time `json:"started_at"`
	Actions   []string  `json:"actions"`
	mu        sync.Mutex
}

func NewState() *State {
	return &State{
		StartedAt: time.Now(),
		Actions:   []string{},
	}
}

func (s *State) AddAction(action string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Actions = append(s.Actions, action)
}

func (s *State) Save(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
