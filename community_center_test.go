package rocket_test

import (
	"testing"

	"rocket"
)

func TestNumberOfClients(t *testing.T) {
	cc := rocket.NewCommunityCenter()
	expected := 5
	for i := 0; i < expected; i++ {
		cc.Register(&rocket.Client{})
	}

	got := cc.NumberOfClients()
	if got != expected {
		t.Errorf("expected number of client %d, got %d", expected, got)
	}
}
