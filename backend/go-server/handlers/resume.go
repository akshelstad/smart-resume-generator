package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResumeRequest struct {
	Name       string `json:"name"`
	Experience string `json:"experience"`
}

type AIResponse struct {
	Resume string `json:"resume"`
}

func GenerateResume(c *gin.Context) {
	var req ResumeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aiURL := "http://localhost:5001/generate-resume"
	jsonData, _ := json.Marshal(req)

	resp, err := http.Post(aiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to communicate with AI service"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var aiResp AIResponse
	json.Unmarshal(body, &aiResp)

	c.JSON(http.StatusOK, gin.H{"resume": aiResp.Resume})
}
