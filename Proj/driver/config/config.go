package config

import "elevator/config"

type Node struct {
	Id        string
	Available bool
	Elevator  config.Elevator
}

var KnownNodes = make([]Node, 0)

func NewNode(id string, e config.Elevator) Node {
	var n Node

	return n
}
