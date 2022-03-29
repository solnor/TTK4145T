package assigner

import (
	"driver/config"
	"elevator/cost"
)

//Calculation of best cost-effective elevator
//TODO: Cab orders?
func assignOrderByCost(elevators []*config.Node) { //orders elevio.ButtonEvent) {
	minimumCost := 100000
	var elevatorCost int64
	var calculatedElevator *config.Node

	for _, e := range elevators {
		elevatorCost = cost.TimeToIdle(e)
		if elevatorCost < minimumCost && config.Available {
			minimumCost = elevatorCost
			calculatedElevator = e
		}
		//failcheck?
	}
	return calculatedElevator

}
