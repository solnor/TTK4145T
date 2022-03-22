package fmt

import (
	. "elevator/config"
	. "elevator/elevator"
	. "elevator/elevio"
	. "elevator/requests"
	. "elevator/timer"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

var Elevator1 Elevator

func Fsm_init() {
	Elevator1 = NewElevator()
}

// var DoorTimer = time.NewTimer(time.Duration(3 * time.Second))
var Obstruction bool = false
var PrevDir MotorDirection

func GetFunctionname(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func setAllLights(es Elevator) {
	for floor := 0; floor < N_FLOORS; floor++ {
		//for brn := range ButtonType { //:= 0; btn < N_BUTTONS; btn++ {
		var btn ButtonType
		for btn = BT_HallUp; btn <= BT_Cab; btn++ {
			if es.Requests[floor][btn] == 1 {
				SetButtonLamp(btn, floor, true)
			} else {
				SetButtonLamp(btn, floor, false)
			}
		}

	}
}

func Fsm_onRequestButtonPress(btn_floor int, btn_type ButtonType) {
	fmt.Printf("\n\n%s(%d, %v)\n", GetFunctionname(Fsm_onRequestButtonPress), btn_floor, btn_type)
	Elevator_print(Elevator1)

	switch Elevator1.Behaviour {
	case EB_DoorOpen:
		if Requests_shouldClearImmediately(Elevator1, btn_floor, btn_type) {
			//timer_start.reset(elevator.Config.DoorOpenDuration_s)
			DoorTimer.Reset(3 * time.Second)

		} else {
			Elevator1.Requests[btn_floor][btn_type] = 1
		}
		break
	case EB_Moving:
		Elevator1.Requests[btn_floor][btn_type] = 1
		break
	case EB_Idle:
		Elevator1.Requests[btn_floor][btn_type] = 1
		var a Action
		a = a.Requests_nextAction(Elevator1)
		Elevator1.Dirn = a.Dirn
		Elevator1.Behaviour = a.Behaviour
		switch a.Behaviour {
		case EB_DoorOpen:
			SetDoorOpenLamp(true)
			DoorTimer.Reset(3 * time.Second)
			Elevator1 = Requests_clearAtCurrentFloor(Elevator1)
			break
		case EB_Moving:
			// if !Obstruction
			SetMotorDirection(Elevator1.Dirn)
			// fmt.Printf("Set motor direction to %v\n", Elevator1.Dirn)
			break
		case EB_Idle:
			break
		}
		break
	}

	setAllLights(Elevator1)
	fmt.Printf("\nNew State: \n")
	Elevator_print(Elevator1)
}

func Fsm_onFloorArrival(newFloor int) {
	fmt.Printf("\n\n%s(%d)\n\n", GetFunctionname(Fsm_onFloorArrival), newFloor)
	Elevator_print(Elevator1)
	Elevator1.Floor = newFloor
	SetFloorIndicator(Elevator1.Floor)

	switch Elevator1.Behaviour {
	case EB_Moving:
		if Requests_shouldStop(Elevator1) {
			Elevator1.Dirn = MD_Stop
			SetMotorDirection(Elevator1.Dirn)
			SetDoorOpenLamp(true)
			Elevator1 = Requests_clearAtCurrentFloor(Elevator1)
			//timer_start(elevator.Config.DoorOpenDuration_s)
			DoorTimer.Reset(3 * time.Second)
			setAllLights(Elevator1)
			Elevator1.Behaviour = EB_DoorOpen
		}
		break
	default:
		break
	}
	fmt.Printf("\nNew State: \n")
	Elevator_print(Elevator1)
}

func Fsm_onDoorTimeout() {
	fmt.Printf("\n\n%s()\n\n", GetFunctionname(Fsm_onDoorTimeout))
	Elevator_print(Elevator1)

	switch Elevator1.Behaviour {
	case EB_DoorOpen:
		var a Action
		a = a.Requests_nextAction(Elevator1)
		Elevator1.Dirn = a.Dirn
		Elevator1.Behaviour = a.Behaviour
		switch Elevator1.Behaviour {
		case EB_DoorOpen:
			//timer_start(elevator.Config.DoorOpenDuration_s)
			DoorTimer.Reset(3 * time.Second)
			fmt.Printf("Timer finished3")

			Elevator1 = Requests_clearAtCurrentFloor(Elevator1)
			setAllLights(Elevator1)
			break
		case EB_Moving:
			SetDoorOpenLamp(false)
			if Obstruction {
				PrevDir = Elevator1.Dirn
			} else {
				SetMotorDirection(Elevator1.Dirn)
			}
		case EB_Idle:
			SetDoorOpenLamp(false)
			if Obstruction {
				PrevDir = Elevator1.Dirn
			} else {
				SetMotorDirection(Elevator1.Dirn)
			}
			break
		}

		break
	default:
		break
	}

	fmt.Printf("\nNew state: \n")
	Elevator_print(Elevator1)
}
