package main

import (
	"bytes"
	"dingding/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	apiURL = config.APIURL
	apiKey = config.APIKey
)

type ChatCompletionRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type ChatCompletionResponse struct {
	Choices []ChatChoice `json:"choices"`
}

type ChatChoice struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

func GetChoiceMessage(userMessage string) (string, error) {
	systemMessage := config.SystemMessage
	requestBody := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "system", Content: systemMessage},
			{Role: "user", Content: userMessage},
		},
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling request body: %v", err)
	}

	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", fmt.Errorf("error creating HTTP request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var completionResponse ChatCompletionResponse
	err = json.Unmarshal(responseBody, &completionResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response body: %v", err)
	}

	if len(completionResponse.Choices) == 0 {
		return "", fmt.Errorf("no choices returned")
	}

	return completionResponse.Choices[0].Message.Content, nil
}
