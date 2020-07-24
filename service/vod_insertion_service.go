package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//These probably don't belong here? Maybe a models folder or something
type replacer interface {
	replace()
}

type videoReplacer struct{}

type audioReplacer struct{}

func (ar audioReplacer) replace() {
	//Impl
}

func (vr videoReplacer) replace() {
	//Impl
}

func fetchReplacer(fileExtension string) (replacer, error) {
	switch fileExtension {
	case "#EXT-X-STREAM-INF":
		return videoReplacer{}, nil
	case "#EXT-X-MEDIA:TYPE=AUDIO":
		return audioReplacer{}, nil
	default:
		return nil, errors.New("Cannot find valid replacer type")
	}
}

/*ReplacePlaylistWithServerEndpoints is a method that will take in a master playlist as i
 *input and replace each subplaylist with an endpoint to be called by the browser
 */
func ReplacePlaylistWithServerEndpoints(playlistURL string, baseURL string) (string, error) {
	manifest, err := getManifestFromResponse(playlistURL)

	if err != nil {
		return "", err
	}
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
