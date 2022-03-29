package config

import "elevator/config"

type Node struct {
	Id        string
	Connected bool
	Elevator  config.Elevator
}

var KnownNodes = make([]Node, 0)
