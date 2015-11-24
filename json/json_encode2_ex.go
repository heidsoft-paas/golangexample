package main

import (
	"encoding/json"
	"os"
)

type Object interface {}

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
	group := &ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		Valid:  true,
		Object: &Inner{
			Age: 11,
		},
	}
	en := json.NewEncoder(os.Stdout)
	en.Encode(group)
}

//output:
//{"id":1,"name":"Reds","colors":["Crimson","Red","Ruby","Maroon"],"valid":true,"Inner":{"age":11}}

