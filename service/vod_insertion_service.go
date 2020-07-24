package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

/*ReplacePlaylistWithServerEndpoints is a method that will take in a master playlist as i
 *input and replace each subplaylist with an endpoint to be called by the browser
 */
func ReplacePlaylistWithServerEndpoints(playlistURL string, baseURL string) (string, error) {

	const videoStreamRegex = "#EXT-X-STREAM-INF"
	const audioStreamRegex = "#EXT-X-MEDIA:TYPE=AUDIO"

	manifest, err := getManifestFromResponse(playlistURL)

	if err != nil {
		return "", err
	}

	manifestSlice := strings.Split(manifest, "\n")

	for _, manifestChunk := range manifestSlice {
		videoRegexMatched, err := regexp.Match(videoStreamRegex, []byte(manifestChunk))
		if err != nil {
			fmt.Println("Video regex error: ", err)
		}

		fmt.Println("Video Regex Matched: ", videoRegexMatched)
		audioRegexMatched, err := regexp.Match(audioStreamRegex, []byte(manifestChunk))
		if err != nil {
			fmt.Println("Audio regex error: ", err)
		}
		fmt.Println("Audio Regex Matched: ", audioRegexMatched)
		fmt.Println(manifestChunk)
	}

	fmt.Println(reflect.TypeOf(manifestSlice))
	// fmt.Println(manifestSlice)

	return manifest, nil

}

func getManifestFromResponse(playlistURL string) (string, error) {
	resp, err := http.Get(playlistURL)
	if err != nil {
		fmt.Println("Playlist could not be found")
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body")
		return "", err
	}
	return string(body), nil
}
