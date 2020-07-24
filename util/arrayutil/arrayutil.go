package util

import "strings"

//A set of ES6 style array methods taken from https://gobyexample.com/collection-functions

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

//Map returns a new slice based on function supplied to it
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

//DestructureKeyValuePair will return a string that is split as a key value pair in the format (key, value)
func DestructureKeyValuePair(stringToSplit string, separator string) (string, string) {
	x := strings.Split(stringToSplit, separator)
	return x[0], x[1]
}
