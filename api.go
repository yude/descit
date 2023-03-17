package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
)

var messages []Message
const api_endpoint = "https://api.openai.com/v1/chat/completions"

func getResponse(apiKey string) OpenaiResponse {
	requestBody := OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	req_json, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", api_endpoint, bytes.NewBuffer(req_json))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		println("Error: ", err.Error())
		return OpenaiResponse{}
	}

	messages = append(messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Messages.Content,
	})

	return response
}