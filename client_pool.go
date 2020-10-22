package rocket

type ClientPool interface {
	Broadcast(msg []byte)
	NumberOfClients() int
	Register(Clienter) error
	Unregister(Clienter)
}
