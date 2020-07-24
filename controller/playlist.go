package controller

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/service"
)

//RegisterControllers is a method that will register the available controller endpoints with the application
//Whether this is the best way to do controllers... idk man I started learning Go 2 days ago
func RegisterControllers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Fallback")
	})

	/**
	* If a master playlist is present (which it most likely will be, replace all sublplaylists with express endpoints),
	* will also call generate dynamic playlist endpoint
	 */
	http.HandleFunc("/generate_master_playlist", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET params were:", r.URL.Query())
		//Full URL of the master playlist that will have the ad inserted at the start
		masterPlaylist := r.URL.Query().Get("masterPlaylist")
		//Base url to append to TS files, should be able to work this out programatically
		baseURL := r.URL.Query().Get("baseUrl")
		if masterPlaylist != "" && baseURL != "" {
			fmt.Println("Both parameters are defined")
			fmt.Println("Master Playlist: ", masterPlaylist)
			fmt.Println("baseURL: ", baseURL)
			service.ReplacePlaylistWithServerEndpoints(masterPlaylist, baseURL)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// fmt.Println("Manifest is: ", manifest)
		}

		fmt.Fprintf(w, "Return master playlist")
	})

	/**
	* If for whatever reason there isn't a master playlist for a particular HLS stream, this endpoint can be used to replace all calls to TS files with the full path
	 */
	http.HandleFunc("/generate_dynamic_playlist", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET params were:", r.URL.Query())
		//Odds are the playlists will be split into audio and video, so will need to know which lines to replace
		format := r.URL.Query().Get("format")
		//URL of the actual bitrate file that will be played
		subPlaylistURL := r.URL.Query().Get("subPlaylistURL")
		if format != "" && subPlaylistURL != "" {
			fmt.Println("Both parameters are defined")
			fmt.Println("Format: ", format)
			fmt.Println("Subplaylist URL", subPlaylistURL)
		}

		fmt.Fprintf(w, "Return dynamic playlist")
	})

	fmt.Println("Listening on port 8080 Container Port 7010 Host")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
