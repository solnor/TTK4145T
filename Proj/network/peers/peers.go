package peers

import (
	"fmt"
	"net"
	"network/conn"
	"sort"
	"time"
)

type PeerUpdate struct {
	Peers []string
	New   string
	Lost  []string
}

const interval = 15 * time.Millisecond
const timeout = 500 * time.Millisecond

func Transmitter(port int, id string, transmitEnable <-chan bool) {

	conn := conn.DialBroadcastUDP(port)
	addr, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("255.255.255.255:%d", port))

	enable := true
	for {
		select {
		case enable = <-transmitEnable:
		case <-time.After(interval):
		}
		if enable {
			conn.WriteTo([]byte(id), addr)
		}
	}
}

func Receiver(port int, peerUpdateCh chan<- PeerUpdate) {

	var buf [1024]byte
	var p PeerUpdate
	lastSeen := make(map[string]time.Time)

	conn := conn.DialBroadcastUDP(port)

	for {
		updated := false

		conn.SetReadDeadline(time.Now().Add(interval))
		n, _, _ := conn.ReadFrom(buf[0:])

		id := string(buf[:n])

		// Adding new connection
		p.New = ""
		if id != "" {
			if _, idExists := lastSeen[id]; !idExists {
				p.New = id
				updated = true
			}

			lastSeen[id] = time.Now()
		}

		// Removing dead connection
		p.Lost = make([]string, 0)
		for k, v := range lastSeen {
			if time.Now().Sub(v) > timeout {
				updated = true
				p.Lost = append(p.Lost, k)
				delete(lastSeen, k)
			}
		}

		// Sending update
		if updated {
			p.Peers = make([]string, 0, len(lastSeen))

			for k, _ := range lastSeen {
				p.Peers = append(p.Peers, k)
			}

			sort.Strings(p.Peers)
			sort.Strings(p.Lost)
			peerUpdateCh <- p
		}
	}
}

// func onNewNode(node PeerUpdate) {
// 	newNodeIsKnown := false
// 	for _, peer := range node.Peers {
// 		if node.New == node.Peers {
// 			newNodeIsKnown = true
// 		}
// 	}
// 	if newNodeIsKnown {
// 		*e, error  = getNodeWithId(node.New)
// 		if error != nil {
// 			fmt.Printf("Error: Could not find elevator with id %s", node.New)
// 			// In case ID is known, but no elevator is associated with the id: Create new node with ID
// 			e = createNewNode(node.New)
// 			knownNodes = append(knownNodes, e)
// 		}
// 		// e.undefined = setNodeDataUndefined(e)
// 		e.connected = true
// 		e.defined = true
// 	} else {
// 		e = createNewNode(node.New)
// 		knownNodes = append(knownNodes, e)
// 	}
// 	// if node.New not in node.Peers {

// 	// } else {

// 	// }
// }
