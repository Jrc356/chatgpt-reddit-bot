package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Role string

const (
	User     Role = "user"
	System   Role = "system"
	endpoint      = "https://api.openai.com/v1/chat/completions"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type requestBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type responseBody struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type GenerateTextParams struct {
	ApiKey      string
	Messages    []Message
	Temperature float64
}

func GenerateText(params GenerateTextParams) (string, error) {
	requestBody := requestBody{
		Model:       "gpt-4-turbo",
		Messages:    params.Messages,
		Temperature: params.Temperature,
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+params.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		return "", fmt.Errorf("bad status from openai %d", resp.StatusCode)
	}

	var responseBody responseBody
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", err
	}

	if len(responseBody.Choices) > 0 {
		return responseBody.Choices[0].Message.Content, nil
	} else {
		return "", errors.New("no text generated")
	}
}
