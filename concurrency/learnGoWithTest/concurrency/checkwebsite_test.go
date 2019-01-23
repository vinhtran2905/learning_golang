package concurrency

import (
	"reflect"
	"testing"
)

func mockWebSiteChecker(url string) bool {
	if url == "s3://aws.com" {
		return false
	}
	return true
}

func TestCheckWebSite(t *testing.T) {
	websites := []string{
		"http://comcast.com",
		"https://qa.comcast.com",
		"s3://aws.com",
	}

	want := map[string]bool{
		"http://comcast.com":     true,
		"https://qa.comcast.com": true,
		"s3://aws.com":           false,
	}
	got := CheckWebSites(mockWebSiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}

}
