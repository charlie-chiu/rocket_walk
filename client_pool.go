package rocket

type ClientPool interface {
	Broadcast(msg []byte)
	NumberOfClients() int
	Register(*Client) error
	Unregister(*Client)
}
