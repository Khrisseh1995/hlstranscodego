package util

import (
	"fmt"
	util "rest_api/util/arrayutil"
	"strings"
)

//FetchValueFromManifestMetadata is a method that will pull out a key value pair metadata tag
func FetchValueFromManifestMetadata(manifestChunk string, matchValue string) {
	splitManifestChunk := strings.Split(manifestChunk, ",")

	filteredString := util.StringFilter(splitManifestChunk, func(s string) bool {
		return strings.Contains(s, matchValue)
	})

	fmt.Println(filteredString)

	// splitKeyValuePair := strings.Split(filteredString)

}
