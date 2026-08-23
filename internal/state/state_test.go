package state_test

import (
	"goflagsmith/internal/state"
	"testing"
)

func TestState_Transition(t *testing.T) {

	s := state.NewState()

	initial := s.Snapshot()
	if initial.Clients || initial.Features {
		t.Fatalf("expected initial state to be false/false, got clients=%v, features=%v", initial.Clients, initial.Features)
	}

	s.SetClientsReady()
	afterClients := s.Snapshot()
	if !afterClients.Clients || afterClients.Features {
		t.Errorf("expected clients=true, features=false, got clients=%v, features=%v", afterClients.Clients, afterClients.Features)
	}

	s.SetFeaturesReady()
	afterFeatures := s.Snapshot()
	if !afterFeatures.Clients || !afterFeatures.Features {
		t.Errorf("expected both true, got clients=%v, features=%v", afterFeatures.Clients, afterFeatures.Features)
	}
}
