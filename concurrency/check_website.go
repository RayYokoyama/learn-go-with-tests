package concurrency

import "time"

func CheckWebsite(url string) bool {
	// request to http
	time.Sleep(3 * time.Second)
	return true
}
