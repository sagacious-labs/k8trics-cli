package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// ConstructURL takes in URL segments and joins them together
//
// If the first segment doesn't starts with http then it is
// auto added to the segment
func ConstructURL(segs ...string) string {
	if len(segs) == 0 {
		return ""
	}

	curl := segs[0]
	if !strings.HasPrefix(curl, "http") {
		curl = fmt.Sprintf("http://%s", curl)
	}

	segs = segs[1:]

	for _, seg := range segs {
		curl = fmt.Sprintf("%s/%s", curl, seg)
	}

	u, err := url.Parse(curl)
	if err != nil {
		return ""
	}

	return u.String()
}

// ContainsString takes a string slice and the value which
// needs to be searched inside the slice, if the value exists
// the function returns true or else false
func ContainsString(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}
