package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Payload struct {
	UserId         string            `json:"userId"`
	UserAttributes map[string]string `json:"userAttributes"`
}

func main() {
	ticker := time.NewTicker(50 * time.Millisecond)

	for range ticker.C {
		data := Payload{
			UserId: "test-user",
			UserAttributes: map[string]string{
				"attr1": "sample-attribute-1",
				"attr2": "sample-attribute-2",
			},
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		body := bytes.NewReader(payloadBytes)

		req, err := http.NewRequest("POST", "http://localhost:4567/decide", body)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("unexpected status code: %d", resp.StatusCode)
		}
	}
}
