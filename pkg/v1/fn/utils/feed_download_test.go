package utils_test

import (
	"encoding/json"
	"fmt"
	"github.com/MickStanciu/go-fn/pkg/v1/fn/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPaginationProcessor(t *testing.T) {
	type feedExample struct {
		NextPageURL string          `json:"next_page_url"`
		Data        json.RawMessage `json:"data,omitempty"`
	}

	url := "https://www.site.fake"

	tests := map[string]struct {
		urlProc     utils.UrlProcessor
		contentProc utils.ContentProcessor
		expectedErr error
	}{
		"when url processor and content processor not throwing errors": {
			urlProc: func(url string) ([]byte, error) {
				urlMap := map[string][]byte{
					"https://www.site.fake":        []byte(`{"next_page_url":"page=1","data":null}`),
					"https://www.site.fake?page=1": []byte(`{"next_page_url":"page=2","data":null}`),
					"https://www.site.fake?page=2": []byte(`{"next_page_url":"","data":null}`),
				}

				v, ok := urlMap[url]
				if !ok {
					return nil, fmt.Errorf("the url %q is not present in the urlMap", url)
				}
				return v, nil
			},
			contentProc: func(content []byte) (string, error) {
				const baseURL = "https://www.site.fake"
				var f feedExample
				err := json.Unmarshal(content, &f)
				if err != nil {
					return "", err
				}
				if f.NextPageURL == "" {
					return "", nil
				}
				return fmt.Sprintf("%s?%s", baseURL, f.NextPageURL), nil
			},
			expectedErr: nil,
		},
		"when url processor errors out": {
			urlProc: func(url string) ([]byte, error) {
				return nil, fmt.Errorf("url processor error")
			},
			contentProc: func(content []byte) (string, error) {
				return "", nil
			},
			expectedErr: fmt.Errorf("url https://www.site.fake triggered a url processor error :url processor error"),
		},
		"when content processor errors out": {
			urlProc: func(url string) ([]byte, error) {
				return []byte(`{"next_page_url":"","data":null}`), nil
			},
			contentProc: func(content []byte) (string, error) {
				return "", fmt.Errorf("content processor error")
			},
			expectedErr: fmt.Errorf("url https://www.site.fake triggered a content processor error :content processor error"),
		},
		"when there is an error during an iteration": {
			urlProc: func(url string) ([]byte, error) {
				urlMap := map[string][]byte{
					"https://www.site.fake":        []byte(`{"next_page_url":"page=1","data":null}`),
					"https://www.site.fake?page=1": []byte(`{"next_page_url":"page=2","data":null}`),
				}

				v, ok := urlMap[url]
				if !ok {
					return nil, fmt.Errorf("the url %q is not present in the urlMap", url)
				}
				return v, nil
			},
			contentProc: func(content []byte) (string, error) {
				const baseURL = "https://www.site.fake"
				var f feedExample
				err := json.Unmarshal(content, &f)
				if err != nil {
					return "", err
				}
				if f.NextPageURL == "" {
					return "", nil
				}
				return fmt.Sprintf("%s?%s", baseURL, f.NextPageURL), nil
			},
			expectedErr: fmt.Errorf("url https://www.site.fake?page=2 triggered a url processor error :the url \"https://www.site.fake?page=2\" is not present in the urlMap"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			e := utils.PaginationProcessor(url, test.urlProc, test.contentProc)
			if e != nil && test.expectedErr != nil {
				require.EqualError(t, e, test.expectedErr.Error())
			} else {
				require.NoError(t, e)
				require.Nil(t, test.expectedErr)
			}
		})
	}

}
