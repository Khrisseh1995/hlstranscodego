package util

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
