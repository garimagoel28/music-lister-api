// song.go
package main

import (
	"encoding/json"
	"net/http"
)

// Song represents a song in a playlist
type Song struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Composers string `json:"composers"`
	MusicURL  string `json:"musicURL"`
}
func getSongDetail(w http.ResponseWriter, r *http.Request) {
	// Extract user ID, playlist ID, and song ID from request
	userID := r.FormValue("userID")
	playlistID := r.FormValue("playlistID")
	songID := r.FormValue("songID")

	// Check if the user with the provided ID exists
	mu.Lock()
	defer mu.Unlock()
	user, found := users[userID]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Find the playlist in the user's playlists
	var playlist *Playlist
	for i := range user.Playlists {
		if user.Playlists[i].ID == playlistID {
			playlist = &user.Playlists[i]
			break
		}
	}

	if playlist == nil {
		http.Error(w, "Playlist not found", http.StatusNotFound)
		return
	}

	// Find the song in the playlist
	var song *Song
	for i := range playlist.Songs {
		if playlist.Songs[i].ID == songID {
			song = &playlist.Songs[i]
			break
		}
	}

	if song == nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	// Return details of the song
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}
