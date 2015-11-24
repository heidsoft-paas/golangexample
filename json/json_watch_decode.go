package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Object interface{}

// EventType defines the possible types of events.
type EventType string

const (
	WSE EventType = "WSE"
)

// Event represents a single event to a watched resource.
type Event struct {
	Type EventType

	Object
}

type WatchState struct {
	Running bool
	Paused  bool
}

type watchEvent struct {
	// The type of the watch event; added, modified, or deleted.
	Type EventType
	Raw  json.RawMessage
}

func main() {

	var (
		we watchEvent
		o  Object
	)
	data := []byte(`{"Type":"WSE","Raw":{"Running":true,"Paused":false}}`)
	err := json.Unmarshal(data, &we)
	if err != nil {
		log.Fatalln("error1:", err)
	}
	switch we.Type {
	case "WSE":
		o = new(WatchState)
	}

	fmt.Println(string(we.Raw))
	err = json.Unmarshal(we.Raw, o)
	if err != nil {
		log.Fatalln("error:", err)
	}
	fmt.Println("%v", o)
}
