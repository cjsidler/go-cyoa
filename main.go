package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

func loadAdventure(filename string) Adventure {
	adventureBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var adventure Adventure
	json.Unmarshal(adventureBytes, &adventure)
	return adventure
}

func getAdventureHandler(filename string) http.HandlerFunc {
	adventure := loadAdventure(filename)

	return func(w http.ResponseWriter, r *http.Request) {
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
}

func main() {
	filename := flag.String("filename", "adventure.json", "the JSON file with the Choose-Your-Own-Adventure story")
	flag.Parse()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", getAdventureHandler(*filename))

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
