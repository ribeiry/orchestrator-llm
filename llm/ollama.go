package llm

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	BaseURL string
	Model   string
}

func (c *Client) Generate(prompt string) (string, error) {
	req := map[string]interface{}{
		"model":  c.Model,
		"prompt": prompt,
		"stream": false,
	}

	body, _ := json.Marshal(req)

	resp, err := http.Post(
		c.BaseURL,
		"application/json",
		bytes.NewBuffer(body),
	)
	log.Println(c.BaseURL)
	log.Println(bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		Response string `json:"response"`
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	return res.Response, err
}
