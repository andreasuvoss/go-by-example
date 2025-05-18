package main 

import "fmt"

type ServerState int

const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)

var stateName = map[ServerState]string{
    StateIdle: "idle",
    StateConnected: "connected",
    StateError: "error",
    StateRetrying: "retrying",
    8: "unknown??",
}

func (ss ServerState) String() string {
    return stateName[ss]
}

func main() {
    ns := transistion(StateIdle)
    fmt.Println(ns)

    ns2 := transistion(ns)
    fmt.Println(ns2)

    // This does not compile, however if we dont declare it as a variable but just pass the number it will
    // thisIsAnInteger := 7
    // ns3 := transistion(thisIsAnInteger);
    // fmt.Println(ns3)
}

func transistion(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
