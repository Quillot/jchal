package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Define a struct Page
type Page struct {
	ID	int	`json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
}

// Define a Page.toString function
// Converts a page to JSON
func (p Page) toString() string {
	return toJson(p)
}

// Return the JSON encoding of p
func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

func main() {
	pages := getPages()
	for _, p := range pages {
		fmt.Println(p.toString())
	}

	// fmt.Println(toJson(pages))
}

// Read the pages from JSON
// Returns all the pages
func getPages() []Page {
	raw, err := ioutil.ReadFile("pages.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Page
	json.Unmarshal(raw, &c)
	return c
}

