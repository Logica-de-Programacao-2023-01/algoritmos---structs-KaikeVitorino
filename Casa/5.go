package main

import (
	"fmt"
)

type Playlist struct {
	Title string
	Songs []string
}

func findPlaylists(playlists []Playlist, songTitle string) []string {
	var result []string
	for _, playlist := range playlists {
		for _, song := range playlist.Songs {
			if song == songTitle {
				result = append(result, playlist.Title)
				break
			}
		}
	}
	return result
}

func main() {
	playlists := []Playlist{
		{Title: "Rock", Songs: []string{"Bohemian Rhapsody", "Stairway to Heaven"}},
		{Title: "Pop", Songs: []string{"Thriller", "Like a Prayer"}},
		{Title: "Classics", Songs: []string{"Bohemian Rhapsody", "Moonlight Sonata"}},
	}

	fmt.Println(findPlaylists(playlists, "Bohemian Rhapsody"))
}
