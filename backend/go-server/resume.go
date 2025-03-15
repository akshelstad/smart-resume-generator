package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type ResumeRequest struct {
	Name       string `json:"name"`
	Experience string `json:"experience"`
}

type AIResponse struct {
	Resume string `json:"resume"`
}

func (cfg *config) handlerGenerateResume(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "", errors.New("invalid request method"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error reading request body", err)
		return
	}
	defer r.Body.Close()

	var req ResumeRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid JSON format", err)
		return
	}

	client := openai.NewClient(cfg.apiKey)

	prompt := fmt.Sprintf(
		"Generate a professional resume summary for %s with experience in %s.",
		req.Name, req.Experience,
	)

	resp, err := client.CreateChatCompletion(
		r.Context(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a professional resume writer."},
				{Role: "user", Content: prompt},
			},
			MaxTokens: 200,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error communicating with AI service", err)
		return
	}

	respondWithJSON(w, http.StatusOK, resp.Choices[0].Message.Content)
}
