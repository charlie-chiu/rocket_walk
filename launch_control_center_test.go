package rocket_test

import (
	"bytes"
	"sync"
	"testing"
	"time"

	"rocket"
)

type SpyClientPool struct {
	broadMessages [][]byte

	mu sync.Mutex
}

func (p *SpyClientPool) Broadcast(msg []byte) {
	p.mu.Lock()
	p.broadMessages = append(p.broadMessages, msg)
	p.mu.Unlock()
}

func (SpyClientPool) NumberOfClients() int {
	panic("implement me")
}

func (SpyClientPool) Register(client rocket.Clienter) error {
	panic("implement me")
}

func (SpyClientPool) Unregister(client rocket.Clienter) {
	panic("implement me")
}

func TestLaunchControlCenter_Run(t *testing.T) {
	clientPool := SpyClientPool{}
	lcc := rocket.NewLCC(&clientPool)
	go lcc.Run(0)

	waitForProcess()

	expectedMsg := [][]byte{
		[]byte(`{"name":"on_state","payload":{"name":"ready"}}`),
		[]byte(`{"name":"on_state","payload":{"name":"betend"}}`),
		[]byte(`{"name":"on_state","payload":{"name":"launch"}}`),
		[]byte(`{"name":"on_state","payload":{"name":"bust"}}`),
		[]byte(`{"name":"on_state","payload":{"name":"end"}}`),
	}

	clientPool.mu.Lock()
	defer clientPool.mu.Unlock()
	for i, msg := range expectedMsg {
		if i+1 > len(clientPool.broadMessages) {
			t.Fatalf("%dth msg not exsits, expected %s", i+1, string(msg))
		}

		if bytes.Compare(msg, clientPool.broadMessages[i]) != 0 {
			t.Errorf("%dth msg assert fail", i)
			t.Logf("want %s", msg)
			t.Logf(" got %s", clientPool.broadMessages[i])
		}
	}
}

func waitForProcess() {
	time.Sleep(10 * time.Millisecond)
}
