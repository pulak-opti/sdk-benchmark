package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

type Payload struct {
	UserId         string            `json:"userId"`
	UserAttributes map[string]string `json:"userAttributes"`
}

var requestsSent int32
var responsesReceived int32

func main() {
	ticker := time.NewTicker(50 * time.Millisecond)

	// Set up signal catching
	sigs := make(chan os.Signal, 1)

	// Catch all signals since not explicitly listing
	signal.Notify(sigs, os.Interrupt)
	// Method invoked upon seeing signal
	go func() {
		<-sigs
		log.Printf("Requests Sent: %d", requestsSent)
		log.Printf("Responses Received: %d", responsesReceived)
		os.Exit(1)
	}()

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

		req, err := http.NewRequest("POST", "http://127.0.0.1:8000/decide", body)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Optimizely-SDK-Key", "RiowyaMnbPxLa4dPWrDqu")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		resp.Body.Close()

		atomic.AddInt32(&requestsSent, 1)

		if resp.StatusCode == http.StatusOK {
			atomic.AddInt32(&responsesReceived, 1)
		} else {
			log.Fatalf("unexpected status code: %d", resp.StatusCode)
		}
	}
}
