package main

import (
	"encoding/json"
	"fmt"
)

type Restaurant struct {
	Restaurant RestaurantData `json:"restaurant"`
}

type RestaurantData struct {
	Name  string `json:"name"`
	Owner Owner  `json:"owner"`
}

type Owner struct {
	Name string `json:"name"`
}

func main() {
	data := `{"restaurant":{"name":"Tickets","owner":{"name":"Ferran"}}}`
	r := Restaurant{}
	json.Unmarshal([]byte(data), &r)

	fmt.Printf("%+v", r)
}
