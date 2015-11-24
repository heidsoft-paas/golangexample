package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Object interface{}

type Inner struct {
	Age int `json:"age"`
}

type ColorGroup struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Colors []string `json:"colors"`
	Valid  bool     `json:"valid"`
	Object
}

func main() {
	out, in := io.Pipe()
	group := &ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		Valid:  true,
		Object: &Inner{
			Age: 11,
		},
	}
	go func() {
		en := json.NewEncoder(in)
		en.Encode(group)
		in.Close()
	}()

	var v ColorGroup
	done := make(chan struct{})
	go func() {
		de := json.NewDecoder(out)
		de.Decode(&v)
		close(done)
	}()
	<-done
	fmt.Printf("%#v", v)
}
