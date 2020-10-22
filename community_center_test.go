package rocket_test

import (
	"bytes"
	"net/http"
	"testing"

	"rocket"
)

type SpyClient struct {
	writeMessages [][]byte
}

func (SpyClient) ServeWS(w http.ResponseWriter, r *http.Request) error {
	panic("implement me")
}

func (SpyClient) ListenJSON(wsMsg chan []byte) {
	panic("implement me")
}

func (c *SpyClient) WriteMsg(msg []byte) {
	c.writeMessages = append(c.writeMessages, msg)
}

func TestCommunityCenter_Broadcast(t *testing.T) {
	// arrange
	cc := rocket.NewCommunityCenter()
	clientPool := make([]*SpyClient, 0)
	for i := 0; i < 5; i++ {
		spyClient := &SpyClient{}
		cc.Register(spyClient)
		clientPool = append(clientPool, spyClient)
	}
	msg := []byte("broadcast msg")

	// act
	cc.Broadcast(msg)

	// assert
	for i, spyClient := range clientPool {
		if len(spyClient.writeMessages) < 1 {
			t.Fatalf("spy client %d haven't write any messages", i)
		}

		got := spyClient.writeMessages[0]
		if bytes.Compare(got, msg) != 0 {
			t.Errorf("message not equal on client %d, want %s, got %s", i, msg, got)
		}
	}
}

func TestNumberOfClients(t *testing.T) {
	numberOfClient := 5
	cc := rocket.NewCommunityCenter()
	for i := 0; i < numberOfClient; i++ {
		cc.Register(&rocket.Client{})
	}

	got := cc.NumberOfClients()
	if got != numberOfClient {
		t.Errorf("expected number of client %d, got %d", numberOfClient, got)
	}
}
