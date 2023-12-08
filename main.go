// main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/viewProfile", viewProfile)
	http.HandleFunc("/getAllSongsOfPlaylist", getAllSongsOfPlaylist)
	http.HandleFunc("/createPlaylist", createPlaylist)
	http.HandleFunc("/addSongToPlaylist", addSongToPlaylist)
	http.HandleFunc("/deleteSongFromPlaylist", deleteSongFromPlaylist)
	http.HandleFunc("/deletePlaylist", deletePlaylist)
	http.HandleFunc("/getSongDetail", getSongDetail)

	// Start the server
	port := 8080
	fmt.Printf("Server running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
