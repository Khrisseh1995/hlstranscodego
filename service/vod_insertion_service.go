package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*ReplacePlaylistWithServerEndpoints is a method that will take in a master playlist as i
 *input and replace each subplaylist with an endpoint to be called by the browser
 */
func ReplacePlaylistWithServerEndpoints(playlistURL string, baseURL string) (string, error) {

	manifest, err := getManifestFromResponse(playlistURL)

	if err != nil {
		return "", err
	}

	manifestSlice := strings.Split(manifest, "\n")

	for index, manifestChunk := range manifestSlice {

		rp := replaceParams{
			baseURL: baseURL,
			index:   index,
		}

		replacer, err := FetchReplacer(manifestChunk)
		if err == nil {
			fmt.Println(replacer)
			replacer.replace(manifestSlice, rp)
		}
	}

	manifest = strings.Join(manifestSlice, "\n")

	return manifest, nil
}

//ReplaceSubPlaylistWithFullURLs due to files now being served from a different origin, the full URL will now have to be
//spliced into the manifest in order to know the location of the TS files
func ReplaceSubPlaylistWithFullURLs(
	subPlaylistURL string,
	streamData string,
	format string,
	baseURL string,
) (string, error) {
	manifest, err := getManifestFromResponse(subPlaylistURL)

	if err != nil {
		return "", err
	}

	fmt.Println(manifest)

	//CANNY assumption, the files could be aac, fmp4, or other types of format; this will need to change
	const mediaFileExtension = ".ts"

	return "Success!", nil
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
