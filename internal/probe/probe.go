package main

import "net/http"

// probe calls the API server to check what we can do
func probe() error {
	// define http calls here, e.g.: http.Get(defaultServerURL + "my-endpoint")

	for _, u := range []string{
		"https://jobicy.com/api/v2/remote-jobs?count=20&tag=python",
		"https://jobicy.com/api/v2/remote-jobs?count=15&geo=canada",
		"https://jobicy.com/api/v2/remote-jobs?count=30&geo=usa&industry=copywriting",
		"https://jobicy.com/api/v2/remote-jobs?count=10&industry=supporting",
	} {
		if _, err := http.Get(u); err != nil {
			return err
		}
	}

	return nil
}
