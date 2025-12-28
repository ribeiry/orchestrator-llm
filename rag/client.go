package rag

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	BaseURL string
}

type SearchRequest struct {
	Query string `json:"query"`
	TopK  int    `json:"topK"`
}

type Chunk struct {
	Content string `json:"content"`
	Source  string `json:"source"`
}

func (c *Client) Search(query string) ([]Chunk, error) {
	body, _ := json.Marshal(SearchRequest{
		Query: query,
		TopK:  5,
	})

	resp, err := http.Post(
		c.BaseURL+"/rag/search",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err

	}

	defer resp.Body.Close()

	var chunks []Chunk
	err = json.NewDecoder(resp.Body).Decode(&chunks)
	return chunks, err
}
