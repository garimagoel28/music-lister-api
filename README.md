# MusicLister Go API

An HTTP JSON Go API for the MusicLister application, offering a range of routes to facilitate playlist creation, updates, and user management. The API includes the following routes:

## Login

- **Purpose:** Allows users to log in using a unique secret code.
- **Functionality:** Returns user details if the user is found.

## Register

- **Purpose:** Facilitates user registration by providing name and email address.
- **Functionality:** Returns comprehensive details of the newly created user, including their secret code and user ID.

## View Profile

- **Purpose:** Provides a comprehensive view of all current playlists associated with a user.
- **Functionality:** Returns details of each playlist, encompassing all user attributes.

## Get All Songs of Playlist

- **Purpose:** Retrieves all songs within a specific playlist.
- **Functionality:** Returns details of each song, including all song attributes.

## Create Playlist

- **Purpose:** Enables the creation of a new playlist with user-defined attributes.
- **Functionality:** Allows users to specify various attributes for the new playlist.

## Add Song to Playlist

- **Purpose:** Facilitates the addition of a new song to an existing playlist.
- **Functionality:** Updates the playlist by appending the new song with all its associated attributes.

## Delete Song from Playlist

- **Purpose:** Permits the removal of a song from a playlist.
- **Functionality:** Updates the playlist by excluding the specified song.

## Delete Playlist

- **Purpose:** Allows users to delete a playlist they no longer need.
- **Functionality:** Removes the specified playlist from the user's profile.

## Get Song Detail

- **Purpose:** Retrieves all attributes of a particular song.
- **Functionality:** Returns comprehensive details of the specified song.

This API empowers MusicLister users to seamlessly manage their playlists, from creation and updates to user authentication and profile viewing.
