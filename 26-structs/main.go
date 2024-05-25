package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func basicStruct() {
	type Movie struct {
		Title  string
		Genre  string
		Rating int
	}

	movie := Movie{Title: "a", Genre: "b", Rating: 10}
	fmt.Println(movie)
	fmt.Printf("movie: %+v\n", movie)
	fmt.Printf("movie title: %s\n", movie.Title)
}

func nestedStruct() {
	type song struct {
		title, artist string
	}
	type playlist struct {
		title string
		songs []song
	}

	song1 := song{title: "hello", artist: "a"}
	song2 := song{title: "hello", artist: "b"}
	songs := []song{song1, song2}
	pl := playlist{title: "myplaylist", songs: songs}
	fmt.Printf("playlist: %v\n", pl)

	//cloned value doesn't affect struct
	song3 := pl.songs[0]
	song3.title = "live forever"
	fmt.Printf("playlist: %v\n", pl)

	//need to edit the actual value
	pl.songs[0].title = "live forever"
	fmt.Printf("playlist: %v\n", pl)

	//iterating through a struct
	for index, val := range pl.songs {
		fmt.Print(index)
		fmt.Println(val)
	}

	// map by title
	byTitle := make(map[string]song)
	for _, g := range pl.songs {
		byTitle[g.title] = g
	}
}

func encoding() {
	type permissions map[string]bool

	type user struct {
		Name        string      `json:"username"`
		Password    string      `json:"-"`
		Permissions permissions `json:"perms,omitempty"`
	}
	users := []user{ // #2
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "66", permissions{"write": true}},
	}

	out, err := json.Marshal(users)
	// out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}

func decoding() {
	type user struct {
		Name        string          `json:"username"`
		Permissions map[string]bool `json:"perms"`

		Devices []struct {
			Name    string `json:"name"`
			Battery int    `json:"battery"`
		} `json:"devices"`
	}

	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write.")
		}

		for _, device := range user.Devices {
			fmt.Printf("\n\t[ %-10s: %d%% ]", device.Name, device.Battery)
		}
		fmt.Println()
	}
}

func main() {
	basicStruct()
	nestedStruct()
	encoding()
	decoding()
}
