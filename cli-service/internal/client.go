package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseURLUserManagement = "http://localhost:5001"
	BaseURLMessages       = "http://localhost:5002"
	BaseURLFeed           = "http://localhost:5003"
)

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

// RegisterUser registers a new user and returns the response and error
func RegisterUser(username string) (*http.Response, error) {
	url := BaseURLUserManagement + "/register"
	payload := map[string]string{"username": username}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %w", err)
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %w", err)
	}

	// Ensure we return a proper error if response is nil
	if response == nil {
		return nil, fmt.Errorf("no response received from the server")
	}

	return response, nil
}

// CreateMessage sends a message and returns the response and error
func CreateMessage(username, content string) (*http.Response, error) {
	url := BaseURLMessages + "/messages"
	payload := map[string]string{"username": username, "content": content}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %w", err)
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %w", err)
	}

	if response == nil {
		return nil, fmt.Errorf("no response received from the server")
	}

	return response, nil
}

// LikeMessage sends a like to a message and returns the response and error
func LikeMessage(messageID string) (*http.Response, error) {
	url := BaseURLMessages + "/messages/" + messageID + "/like"
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating POST request: %w", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %w", err)
	}

	if response == nil {
		return nil, fmt.Errorf("no response received from the server")
	}

	return response, nil
}

// GetFeed retrieves the latest feed of messages and handles any potential errors
func GetFeed() ([]Message, error) {
	url := BaseURLFeed + "/feed"
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("failed to fetch feed: status code %d, response %s", response.StatusCode, string(body))
	}

	var messages []Message
	if err := json.NewDecoder(response.Body).Decode(&messages); err != nil {
		return nil, fmt.Errorf("error decoding feed response: %w", err)
	}

	return messages, nil
}
