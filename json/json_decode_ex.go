package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock.", "Valid": true}
		{"Name": "Ed", "Text": "Go fmt yourself!", "Valid": true}
		{"Name": "Ed", "Valid": true}
	`
	type Message struct {
		Name, Text string
        Valid bool `json:"Valid"`
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", m)
		//fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
