package groq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Llama3 struct {
	client      *http.Client
	model       string
	userPrompt  string
	sytemPrompt string
}

func newLlama3() *Llama3 {
	return &Llama3{
		client: http.DefaultClient,
		model:  LLAMA3_8B,
	}
}

func (l *Llama3) UserPrompt(prompt string) {
	l.userPrompt = prompt
}

func (l *Llama3) SystemPrompt(prompt string) {
	l.sytemPrompt = prompt
}

func (l *Llama3) Send() (Response, error) {
	var r Response

	// Create a new HTTP request
	jsonData, err := json.Marshal(l.Request())
	if err != nil {
		return r, err
	}

	httpReq, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return r, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))

	resp, err := l.client.Do(httpReq)
	if err != nil {
		return r, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return r, newError(body)
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return r, err
	}

	return r, nil
}

func (l *Llama3) Request() request {
	return request{
		Model: l.model,
		Messages: []message{
			{
				Role:    "user",
				Content: l.userPrompt,
			},
			{
				Role:    "system",
				Content: l.sytemPrompt,
			},
		},
	}
}
