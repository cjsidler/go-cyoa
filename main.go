package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

var templates = template.Must(template.ParseFiles("./templates/view.html"))
var adventure = loadAdventure("adventure.json")

func loadAdventure(filename string) Adventure {
	adventureBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var adventure Adventure
	json.Unmarshal(adventureBytes, &adventure)
	return adventure
}

func getAdventure(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	var storyArc string
	if path[1] == "" {
		storyArc = "intro"
	} else {
		storyArc = path[1]
	}

	err := templates.ExecuteTemplate(w, "view.html", adventure[storyArc])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", getAdventure)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
