package httpimpl

import (
	"testing"
)

func TestGet(t *testing.T) {

}

func TestParse(t *testing.T) {
	url := "http://example.com/asdf/qwer"
	parsedUrl, _ := parseUrl(url)
	if parsedUrl.scheme != Http || parsedUrl.host != "example.com" || parsedUrl.port != 80 || parsedUrl.path != "/asdf/qwer" {
		t.Fatalf("Incorrectly parsed %s, got %v", url, parsedUrl)
	}

	_, err := parseUrl("https://asdf.com")
	if err == nil {
		t.Fatal("expected error")
	}

	url = "http://asdf.com:123/"
	parsedUrl, _ = parseUrl(url)
	if parsedUrl.host != "asdf.com" || parsedUrl.scheme != Http || parsedUrl.port != 123 || parsedUrl.path != "/" {
		t.Fatalf("Incorrectly parsed %s, got %v", url, parsedUrl)
	}
}