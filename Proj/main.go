package main

import (
	"elevator/config"
	"elevator/elevio"

	// . "elevator/requests"
	"fmt"

	. "elevator/fsm"
	. "elevator/timer"
	// . "../Driver-go/timer"
)

func main() {
	fmt.Printf("Startup\n")

	// numFloors := 4

	elevio.Init("localhost:15657", config.NumFloors)
	Fsm_init()

	// var d config.MotorDirection = config.MD_Up
	//elevio.SetMotorDirection(d)
	// var elev config.Elevator
	// fsm.elevator = elev
	// var Obstructions bool = false
	// var PrevDir config.MotorDirection

	drv_buttons := make(chan config.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)
	go Timesda(DoorTimer)

	// time.Sleep(1000 * time.Millisecond)

	if elevio.GetFloor() == -1 {
		elevio.SetMotorDirection(config.MD_Down)
		Elevator1.Dirn = config.MD_Down
		Elevator1.Behaviour = config.EB_Moving
	}
	/*
		for Elevator1.Floor < 0 {
			select {
			case a := <-drv_floors:
				Elevator1.Floor = a
				fmt.Println("Bjeff")
				elevio.SetMotorDirection(config.MD_Stop)
				Elevator1.Dirn = config.MD_Stop
				Elevator1.Behaviour = config.EB_Idle
			default:
				elevio.SetMotorDirection(config.MD_Down)
				Elevator1.Dirn = config.MD_Down
				Elevator1.Behaviour = config.EB_Moving
			}
		}*/

	for {
		select {
		case a := <-drv_buttons:
			Fsm_onRequestButtonPress(a.Floor, a.Button)
		case a := <-drv_floors:
			Fsm_onFloorArrival(a)

		case a := <-drv_obstr:
			fmt.Printf("%+v\n", a)
			if a {
				Elevator1.Obstruction = true
				// PrevDir = Elevator1.Dirn
				// elevio.SetMotorDirection(config.MD_Stop)
			} else {
				Elevator1.Obstruction = false
				// elevio.SetMotorDirection(PrevDir)
			}

		case a := <-drv_stop:
			fmt.Printf("%+v\n", a)
			for f := 0; f < config.NumFloors; f++ {
				for b := config.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}
		case <-Timer:
			Fsm_onDoorTimeout()
		}
	}
}
