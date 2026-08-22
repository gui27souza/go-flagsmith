package state

import "sync"

type ReadinessStatus struct {
	Clients  bool `json:"clients_ready"`
	Features bool `json:"features_ready"`
}

type State struct {
	mu     sync.RWMutex
	status ReadinessStatus
}

func NewState() *State {
	return &State{}
}

func (s *State) SetClientsReady() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status.Clients = true
}

func (s *State) SetFeaturesReady() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status.Features = true
}

func (s *State) Snapshot() ReadinessStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.status
}
