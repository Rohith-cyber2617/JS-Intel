package crawler

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func DiscoverJS(target string) ([]string, error) {

	var jsFiles []string

	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	baseURL, _ := url.Parse(target)

	doc.Find("script").Each(func(i int, s *goquery.Selection) {

		src, exists := s.Attr("src")
		if !exists || src == "" {
			return
		}

		parsedURL, err := url.Parse(src)
		if err != nil {
			return
		}

		jsURL := baseURL.ResolveReference(parsedURL)

		jsFiles = append(jsFiles, jsURL.String())
	})

	return jsFiles, nil
}
