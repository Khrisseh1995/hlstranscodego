package util

import "strings"

//StringFilter filters a string array based on a conditional passed in via a function
//Might be able to use a empty interface to make array generic?
func StringFilter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

//DestructureKeyValuePair will return a string that is split as a key value pair in the format (key, value)
func DestructureKeyValuePair(stringToSplit string, separator string) (string, string) {
	x := strings.Split(stringToSplit, separator)
	return x[0], x[1]
}
