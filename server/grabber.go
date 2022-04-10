package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getNeighbours(cur string) ([]string, error) {
	var res []string
	resp, err := http.Get("https://en.wikipedia.org/w/api.php?action=query&prop=links&pllimit=max&format=json&titles=" + cur)
	if err != nil {
		return nil, fmt.Errorf("can't get http request: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad http request on %s", cur)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, fmt.Errorf("can't copy from body: %w", err)
	}
	var body map[string]interface{}
	if err = json.Unmarshal(buf.Bytes(), &body); err != nil {
		return nil, fmt.Errorf("can't unmarshal body content: %w", err)
	}
	query := body["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})
	for _, v := range pages {
		page := v.(map[string]interface{})
		switch links := page["links"].(type) {
		case []interface{}:
			for _, link := range links {
				_link := link.(map[string]interface{})
				res = append(res, strings.ReplaceAll(_link["title"].(string), " ", "_"))
			}
		default:
			return nil, fmt.Errorf("can't parse links")
		}

	}
	return res, err
}
