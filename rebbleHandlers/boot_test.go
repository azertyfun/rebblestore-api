package rebbleHandlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"pebble-dev/rebblestore-api/common"

	"github.com/adams-sarah/test2doc/test"
)

var server *test.Server

func TestVersion(t *testing.T) {

	url := fmt.Sprintf("%s/admin/version", server.URL)
	//fmt.Println(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Printf("%v\n", res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), common.Buildversionstring) {
		t.Fatalf("version mismatch, expected %v, got %v", common.Buildversionstring, string(body))
	}
	//fmt.Println(body)
}
func TestVersion2(t *testing.T) {
	url := fmt.Sprintf("%s/admin/version", server.URL)
	r, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("expected 200, got %v", r.StatusCode)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), common.Buildversionstring) {
		t.Fatalf("version mismatch, expected %v, got %s", common.Buildversionstring, string(body))
	}
}

func TestHome(t *testing.T) {
	url := fmt.Sprintf("%s/", server.URL)
	r, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("expected 200, got %v", r.StatusCode)
	}
}

func TestApps(t *testing.T) {
	url := fmt.Sprintf("%s/dev/apps", server.URL)
	r, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if r.StatusCode != 200 && r.StatusCode != 500 {
		t.Fatalf("expected 200 or 500 (depending on wether the database has been built or not), got %v", r.StatusCode)
	}
}

func TestBoot(t *testing.T) {
	url := fmt.Sprintf("%s/boot/ios/v3/1/1?app_version=4.3&store_uri=https%3A%2F%2Fsantoku.adamfourney.com", server.URL)
	r, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if r.StatusCode == 404 {
		t.Fatalf("expected non 404, got %v", r.StatusCode)
	}
}
