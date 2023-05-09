package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type UrlProcessor func(string) ([]byte, error)
type ContentProcessor func([]byte) (string, error)

// PaginationProcessor - takes an url, an urlProcessor function and a pageProcessor function
func PaginationProcessor(url string, urlProc UrlProcessor, contentProc ContentProcessor) error {
	r, err := urlProc(url)
	if err != nil {
		return fmt.Errorf("url %s triggered a url processor error :%w", url, err)
	}

	nextPageURL, err := contentProc(r)
	if err != nil {
		return fmt.Errorf("url %s triggered a content processor error :%w", url, err)
	}

	if nextPageURL != "" {
		if err = PaginationProcessor(nextPageURL, urlProc, contentProc); err != nil {
			return err
		}
	}
	return nil
}

// GetUrlDownload - sample url processor that takes an url string and returns its content as []byte or error
func GetUrlDownload(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return nil, fmt.Errorf("received http status: %d", r.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response: %w", err)
	}

	return buf.Bytes(), nil
}
