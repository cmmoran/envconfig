//go:build appengine || go1.5
// +build appengine go1.5

package envconfig

import "os"

var lookupEnv = func(key string) (string, bool) {
	location, ok := os.LookupEnv(key + "_FILE")
	if ok {
		contents, err := os.ReadFile(location)
		if err != nil {
			return "", false
		}
		return string(contents), true
	}
	return os.LookupEnv(key)
}
