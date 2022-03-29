package assigner

import (
	"driver/config"
	"elevator/cost"
)

//Calculation of best cost-effective elevator
func assignOrderByCost(elevators []*config.Node) { //orders elevio.ButtonEvent) {
	minimumCost := 9999
	var elevatorCost int
	var calculatedElevator *config.Node

	for _, elevator := range elevators {
		elevatorCost = int(cost.TimeToIdle(elevator))
		if elevatorCost < minimumCost {
			minimumCost = elevatorCost
			calculatedElevator = elevator
		}

	}

}
