package service

import (
	"errors"
	"fmt"
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
	fmt.Println("Audio replacer")
}

//Slice passed in so pointer value
func (vr videoReplacer) replace(manifestSlice []string, rp replaceParams) {
	fmt.Println("Index: ", rp.index)
	subPlaylist := manifestSlice[rp.index+1]
	fmt.Println(subPlaylist)
	manifestSlice[rp.index+1] = fmt.Sprintf(`http://localhost:7003/generate_dynamic_playlist?subPlaylistUrl=%s/%s&format=video`, rp.baseURL, subPlaylist)
	// if err != nil {
	// fmt.Println()
	// }
	fmt.Println("Video Player")
}

//FetchReplacer returns a different instance of replacer depending on the piece of metadata that is passed to it
// EXT-X-STREAM-INF denotes a video playlist, #EXT-X-MEDIA:TYPE=AUDIO denotes an audio playlist
func FetchReplacer(manifestSlice string) (Replacer, error) {
	//Better ways of doing this than if statement, refactor at some point
	if !strings.Contains(manifestSlice, "#EXT-X-STREAM-INF") && !strings.Contains(manifestSlice, "#EXT-X-MEDIA:TYPE=AUDIO") {
		return nil, errors.New("Cannot find valid replacer type")
	}

	if strings.Contains(manifestSlice, "#EXT-X-STREAM-INF") {
		return videoReplacer{}, nil
	}

	if strings.Contains(manifestSlice, "#EXT-X-MEDIA:TYPE=AUDIO") {
		return audioReplacer{}, nil
	}

	return nil, errors.New("Cannot find valid replacer type")

}
