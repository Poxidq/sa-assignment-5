package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

type APIClient struct {
	BaseURLUserManagement string
	BaseURLMessages       string
	BaseURLFeed           string
}

func NewAPIClient() *APIClient {
	return &APIClient{
		BaseURLUserManagement: os.Getenv("BASE_URL_USER_MANAGEMENT"),
		BaseURLMessages:       os.Getenv("BASE_URL_MESSAGES"),
		BaseURLFeed:           os.Getenv("BASE_URL_FEED"),
	}
}

func (client *APIClient) RegisterUser(username string) (*http.Response, error) {
	url := fmt.Sprintf("%s/register", client.BaseURLUserManagement)
	payload := map[string]string{"username": username}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *APIClient) CreateMessage(username, content string) (*http.Response, error) {
	url := fmt.Sprintf("%s/messages", client.BaseURLMessages)
	payload := map[string]string{"username": username, "content": content}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *APIClient) GetFeed() ([]Message, error) {
	url := fmt.Sprintf("%s/feed", client.BaseURLFeed)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var messages []Message
	if err := json.NewDecoder(response.Body).Decode(&messages); err != nil {
		return nil, err
	}

	return messages, nil
}
