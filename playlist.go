// playlist.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Playlist represents a playlist for a user
type Playlist struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Songs []Song `json:"songs"`
}

func getAllSongsOfPlaylist(w http.ResponseWriter, r *http.Request) {
	// Extract user ID and playlist ID from request
	userID := r.FormValue("userID")
	playlistID := r.FormValue("playlistID")

	// Check if the user with the provided ID exists
	mu.Lock()
	defer mu.Unlock()
	user, found := users[userID]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Find the playlist in the user's playlists
	var playlist Playlist
	for _, p := range user.Playlists {
		if p.ID == playlistID {
			playlist = p
			break
		}
	}

	// Return songs in the playlist
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlist.Songs)
}

func createPlaylist(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from request
	userID := r.URL.Query().Get("userID")

	// Check if the user with the provided ID exists
	mu.Lock()
	defer mu.Unlock()
	user, found := users[userID]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Parse JSON request body to get playlist details
	var newPlaylist Playlist
	err := json.NewDecoder(r.Body).Decode(&newPlaylist)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate unique ID for the new playlist
	newPlaylist.ID = generateUniqueID()

	// Add the new playlist to the user's playlists
	user.Playlists = append(user.Playlists, newPlaylist)
	fmt.Println(user)
	users[userID]=user 
	// Return the details of the newly created playlist
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPlaylist)
}

func addSongToPlaylist(w http.ResponseWriter, r *http.Request) {
	// Extract user ID and playlist ID from request
	userID := r.FormValue("userID")
	playlistID := r.FormValue("playlistID")
	fmt.Println(userID,playlistID)
	// Check if the user with the provided ID exists
	mu.Lock()
	defer mu.Unlock()
	user, found := users[userID]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	fmt.Println(user)
	// Find the playlist in the user's playlists
	var playlist *Playlist
	var index int
	for i := range user.Playlists {
		if user.Playlists[i].ID == playlistID {
			playlist = &user.Playlists[i]
			index = i
			break
		}
	}

	if playlist == nil {
		http.Error(w, "Playlist not found", http.StatusNotFound)
		return
	}

	// Parse JSON request body to get song details
	var newSong Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate unique ID for the new song
	newSong.ID = generateUniqueID()

	// Add the new song to the playlist
	playlist.Songs = append(playlist.Songs, newSong)
	users[userID].Playlists[index] = *playlist
	// Return the details of the newly created song
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSong)
}

func deleteSongFromPlaylist(w http.ResponseWriter, r *http.Request) {
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
	var index int
	for i := range user.Playlists {
		if user.Playlists[i].ID == playlistID {
			playlist = &user.Playlists[i]
			index = i
			break
		}
	}

	if playlist == nil {
		http.Error(w, "Playlist not found", http.StatusNotFound)
		return
	}

	// Find and remove the song from the playlist
	var updatedSongs []Song
	for _, s := range playlist.Songs {
		if s.ID != songID {
			updatedSongs = append(updatedSongs, s)
		}
	}
	playlist.Songs = updatedSongs
	users[userID].Playlists[index]=*playlist
	w.WriteHeader(http.StatusNoContent)
}

func deletePlaylist(w http.ResponseWriter, r *http.Request) {
	// Extract user ID and playlist ID from request
	userID := r.FormValue("userID")
	playlistID := r.FormValue("playlistID")

	// Check if the user with the provided ID exists
	mu.Lock()
	defer mu.Unlock()
	user, found := users[userID]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Find and remove the playlist from the user's playlists
	var updatedPlaylists []Playlist
	for _, p := range user.Playlists {
		if p.ID != playlistID {
			updatedPlaylists = append(updatedPlaylists, p)
		}
	}
	user.Playlists = updatedPlaylists
	users[userID]=user
	w.WriteHeader(http.StatusNoContent)
}
