package service

import (
	"errors"
	"fmt"
	util "rest_api/util/playlistutil"
	"strings"
)

//Replacer is ... polymorphism might be a bit overkill here, but wanted to try it out in golang
type Replacer interface {
	replace(manifestSlice []string, rp replaceParams)
}

type replaceParams struct {
	manifestSlice []string
	baseURL       string
	index         int
}

type videoReplacer struct{}

type audioReplacer struct{}

//Slice passed in so pointer value no need to be explicit
func (ar audioReplacer) replace(manifestSlice []string, rp replaceParams) {
	originalURIValue := util.FetchValueFromManifestMetadata(manifestSlice[rp.index], "URI=")
	audioTagEndpoint := fmt.Sprintf("http://localhost:8080/generate_dynamic_playlist?subPlaylistUrl=%s/%s&format=audio", rp.baseURL, originalURIValue)
	//Make endpoint environment variable for different dev/live endpoints
	replacedAudioTag := strings.Replace(
		manifestSlice[rp.index],
		originalURIValue,
		audioTagEndpoint,
		-1,
	)

	manifestSlice[rp.index] = replacedAudioTag

}

//Slice passed in so pointer value
func (vr videoReplacer) replace(manifestSlice []string, rp replaceParams) {
	subPlaylist := manifestSlice[rp.index+1]
	//Make endpoint environment variable for different dev/live endpoints
	manifestSlice[rp.index+1] = fmt.Sprintf(`http://localhost:8080/generate_dynamic_playlist?subPlaylistUrl=%s/%s&format=video`, rp.baseURL, subPlaylist)
}

//FetchReplacer returns a different instance of replacer depending on the piece of metadata that is passed to it
// EXT-X-STREAM-INF denotes a video playlist, #EXT-X-MEDIA:TYPE=AUDIO denotes an audio playlist
func FetchReplacer(manifestSlice string) (Replacer, error) {
	switch manifestSliceToCheck := manifestSlice; {
	case strings.Contains(manifestSliceToCheck, "#EXT-X-STREAM-INF"):
		return videoReplacer{}, nil
	case strings.Contains(manifestSliceToCheck, "#EXT-X-MEDIA:TYPE=AUDIO"):
		return audioReplacer{}, nil
	default:
		return nil, errors.New("Cannot find valid replacer type")
	}
}
