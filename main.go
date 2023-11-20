package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kr/pretty"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Adventure map[string]StoryArc

func main() {
	adventureBytes, err := os.ReadFile("adventure.json")
	if err != nil {
		fmt.Printf("error reading adventure: %s\n", err)
	}

	var adventure Adventure
	json.Unmarshal(adventureBytes, &adventure)

	fmt.Println("adventure:")
	pretty.Println(adventure)
}
