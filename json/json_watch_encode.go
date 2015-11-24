package main

import (
	"encoding/json"
	"os"
)

type Object interface{}

// EventType defines the possible types of events.
type EventType string

const (
	WSE    EventType = "WSE"
)

// Event represents a single event to a watched resource.
type Event struct {
	Type EventType

	Object
}

type WatchState struct {
	Running    bool
	Paused     bool
}

type watchEvent struct {
	// The type of the watch event; added, modified, or deleted.
	Type EventType
	Point json.RawMessage
}

func main() {
	e := &Event{
		Type: WSE,
		Object: &WatchState{
			Running: true,
			Paused: false,
		},
	}
	
/*	data, err := json.Marshal(e.Object);
	if err != nil{
		fmt.Println("marshal error")
		return
	}
	
	we := &watchEvent{e.Type, data}*/
	en := json.NewEncoder(os.Stdout)
	en.Encode(e)
}
