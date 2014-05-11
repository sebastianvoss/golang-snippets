package main

import (
	"fmt"
	"github.com/gorilla/schema"
)

type Message struct {
	Name    string
	Body    string
	Version int64
}

func main() {
	values := map[string][]string{
		"Name":    {"Sebastian"},
		"Body":    {"123456789"},
		"Version": {"1"},
	}

	message := new(Message)

	decoder := schema.NewDecoder()
	decoder.Decode(message, values)

	fmt.Printf("Message: %v\n", message)
}
