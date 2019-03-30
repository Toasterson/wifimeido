package networks

type Storage interface {
	SaveNetwork(net Network)
}
