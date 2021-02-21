package util

import (
	util "ad_insertion/util/arrayutil"
	"strings"
)

//FetchValueFromManifestMetadata is a method that will pull out a key value pair metadata tag
func FetchValueFromManifestMetadata(manifestChunk string, matchValue string) string {
	splitManifestChunk := strings.Split(manifestChunk, ",")

	filteredString := util.StringFilter(splitManifestChunk, func(s string) bool {
		return strings.Contains(s, matchValue)
	})[0]

	_, value := util.DestructureKeyValuePair(filteredString, "=")

	value = strings.Trim(value, "\"")

	return value
}
