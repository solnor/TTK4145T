package elevator

import (
	. "elevator/config"
	"fmt"
)

func eb_toString(eb ElevatorBehaviour) string {
	if eb == EB_Idle {
		return "EB_idle"
	} else if eb == EB_DoorOpen {
		return "EB_DoorOpen"
	} else if eb == EB_Moving {
		return "EB_Moving"
	} else {
		return "EB_UNDEFINED"
	}
	// eb == EB_Idle       ? "EB_Idle"         :
	// eb == EB_DoorOpen   ? "EB_DoorOpen"     :
	// eb == EB_Moving     ? "EB_Moving"       :
	//                       "EB_UNDEFINED"    ;
}

func Elevator_print(es Elevator) {
	fmt.Printf("  +--------------------+\n")
	fmt.Printf("  |floor = %-2d          |\n  |dirn  = %-12.12s|\n  |behav = %-12.12s|\n",
		es.Floor,
		elevio_dirn_toString(es.Dirn),
		eb_toString(es.Behaviour))
	fmt.Printf("  +--------------------+\n")
	fmt.Printf("  |  | up  | dn  | cab |\n")
	for f := N_FLOORS - 1; f >= 0; f-- {
		fmt.Printf("  | %d", f)
		var btn ButtonType
		for btn = 0; btn <= BT_Cab; btn++ {
			if (f == N_FLOORS-1 && btn == BT_HallUp) ||
				(f == 0 && btn == BT_HallDown) {
				fmt.Printf("|     ")
			} else {
				if es.Requests[f][btn] == 1 {
					fmt.Printf("|  #  ")
				} else {
					fmt.Printf("|  -  ")
				}
				// fmt.Printf(es.requests[f][btn] ? "|  #  " : "|  -  ");
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +--------------------+\n")
}

func elevio_dirn_toString(d MotorDirection) string {
	if d == MD_Up {
		return "MD_Up"
	} else if d == MD_Down {
		return "MD_Down"
	} else if d == MD_Stop {
		return "MD_Stop"
	} else {
		return "MD_UNDEFINED"
	}

	// return
	//     d == D_Up    ? "D_Up"         :
	//     d == D_Down  ? "D_Down"       :
	//     d == D_Stop  ? "D_Stop"       :
	//                    "D_UNDEFINED"  ;
}

func elevio_button_toString(b ButtonType) string {
	if b == BT_HallUp {
		return "B_HallUp"
	} else if b == BT_HallUp {
		return "B_HallDown"
	} else if b == BT_HallUp {
		return "B_Cab"
	} else {
		return "B_UNDEFINED"
	}

	// return
	//     b == B_HallUp       ? "B_HallUp"        :
	//     b == B_HallDown     ? "B_HallDown"      :
	//     b == B_Cab          ? "B_Cab"           :
	//                           "B_UNDEFINED"     ;
}
