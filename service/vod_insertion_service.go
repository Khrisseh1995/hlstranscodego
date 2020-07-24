package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	util "rest_api/util/arrayutil"
	"strings"
)

/*ReplacePlaylistWithServerEndpoints is a method that will take in a master playlist as i
* input and replace each subplaylist with an endpoint to be called by the browser
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
//spliced into the manifest in order to know the location of the media files
func ReplaceSubPlaylistWithFullURLs(
	subPlaylistURL string,
	format string,
	baseURL string,
) (string, error) {
	manifest, err := getManifestFromResponse(subPlaylistURL)
	manifestArray := strings.Split(manifest, "\n")

	if err != nil {
		return "", err
	}

	//CANNY assumption, the files could be aac, fmp4, or other types of format; this will need to be dynamic
	const mediaFileExtension = ".ts"

	replacedManifestStreams := util.Map(manifestArray, func(stream string) string {
		//We know we're at the end of the metadata, change duration to be the same as the ad that is being inserted
		if strings.Contains(stream, "#EXT-X-TARGETDURATION") {
			//The duration of the test ad is 10 seconds, but this will have to be dynamic when serving different ads
			return fmt.Sprintf("#EXT-X-TARGETDURATION:%s", "10")
		}
		if strings.Contains(stream, mediaFileExtension) {
			//File directories may look like ../../file so need to parse this and pop the corresponding amount off the base
			//Not sure how stable this is tho...
			re := regexp.MustCompile(regexp.QuoteMeta("../")) //Escape special chars ^.^
			fmt.Println(stream)
			backFileMatches := re.FindAllString(stream, -1)
			//If the array is not empty, we will need to pop the corresponding amount of '../' off the full URL
			// '../' denotes going back a file if you didn't already know...
			var fileBackCount int
			if len(backFileMatches) > 0 {
				fileBackCount = len(backFileMatches)
			} else {
				fileBackCount = 0
			}

			fmt.Println("Back file matches: ", backFileMatches)
			fmt.Println("File back count: ", fileBackCount)

		}
		return "s"
	})

	fmt.Println(replacedManifestStreams)

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
