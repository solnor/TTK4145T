package timer

import (
	"time"
)

var Timer = make(chan bool)

// var DoorTimer *time.Timer
//Initialize timers
var DoorTimer = time.NewTimer(time.Duration(1000 * time.Second))

func Timesda(timer *time.Timer) {
	for {
		<-timer.C
		Timer <- true
	}
}

// func StartTimer(seconds int64) {
// 	//Initialize timers
// 	DoorTimer = time.NewTimer(time.Duration(1000 * time.Second))

// }

func ResetTimer(seconds int64) {
	DoorTimer.Reset(3 * time.Second)
}
