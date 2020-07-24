package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
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

	for index, manifestChunk := range manifestSlice {

		rp := replaceParams{
			manifestSlice: manifestSlice,
			baseURL:       baseURL,
			index:         index,
		}

		replacer, err := FetchReplacer(manifestChunk)

		if err != nil {
			replacer.replace(rp)
		}

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
