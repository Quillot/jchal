// From this lovely https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Stall struct {
	Id int	`json:"id"`
	StallName string `json:"stall-name"`
	StallDesc string `json:"stall-desc"`
	Contact string	`json:"contact"`
	Items []Item	`json:"items"`
}

type Item struct {
	ItemName string `json:"item-name"`
	ItemDesc string `json:"item-desc"`
	Price float64
}

func (s Stall) toString() string {
	return toJson(s)
}

// Convert the stall interface to JSON
func toJson(s interface{}) string {
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	return string(bytes)
}

// Read all the stalls from a JSON file
func getStalls() []Stall {
	raw, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshall the read json into s
	var s []Stall
	json.Unmarshal(raw, &s)
	return s
}

func main() {
	stalls := getStalls()	
	for _, s := range stalls {
		fmt.Println(s.toString())
	}
}