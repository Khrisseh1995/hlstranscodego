package service

import (
	"errors"
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
func ReplacePlaylistWithServerEndpoints(manifest string, baseURL string) {
	//Impl
}
