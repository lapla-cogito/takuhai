package track

import (
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func CheckExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func req(endpoint string, val url.Values) (*goquery.Document, error) {
	req, err := http.PostForm(endpoint, val)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		return nil, errors.New("status code is not 200")
	}
	return goquery.NewDocumentFromReader(req.Body)
}
