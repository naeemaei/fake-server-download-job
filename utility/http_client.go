package utility

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

func HttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1,
		},
		Timeout: 10 * time.Second,
	}

	return client
}

func SendRequest(client *http.Client, endpoint string, method string) []byte {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte("Post this data")))
	if err != nil {
		log.Printf("Error Occured. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Couldn't parse response body. %+v", err)
	}
	log.Printf("Endpoint: %s, Download size: %d bytes", endpoint, len(body))
	return body
}
