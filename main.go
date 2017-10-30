package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Stall struct {
	Id int	`json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Contact string	`json:"contact"`
	Items []Item	`json:"items"`
}

type Item struct {
	Name string
	Price float64
	// Image will come from assets folder <Name.png>
	Desc string
}

func makeRoute(path string, output string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(output))
	})
}

func getJson() []Item {
	stall, _ := ioutil.ReadFile("carlos.json")
	var data []Item
	err := json.Unmarshal(stall, &data)
	if err != nil {
		fmt.Println("Cannot unmarshal json", err)
	}
	return data
}

func main() {

	// stall := &Stall{
	// 	Name: "Carlo's Quesadillas",
	// 	Desc: []byte("Good stuff"),
	// 	Contact: "facebook.com/carloque",
	// 	Items: []Item {
	// 		Item {
	// 			Name: "Cheese Quesadilla",
	// 			Desc: []byte("It's very cheesy"),
	// 			Price: 90.00,
	// 		},
	// 		Item {
	// 			Name: "Meat Quesadilla",
	// 			Desc: []byte("It's very meaty"),
	// 			Price: 100.00,
	// 		},
	// 	},
	// }

	// fmt.Println(stall.Name)
	p := getJson()
	fmt.Println(p)
	makeRoute("/", "Welcome to the JSEC Challenge Page")
	http.ListenAndServe(":8080", nil)

}