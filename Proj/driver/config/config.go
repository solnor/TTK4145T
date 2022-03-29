package config

import "elevator/config"

type elevatorData struct {
	id        string
	connected bool
	elevator  config.Elevator
}
