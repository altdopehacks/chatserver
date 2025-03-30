package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"chatserver/server/models"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// HandleAssistant はアシスタントエンドポイントのハンドラーです
func HandleAssistant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// リクエストのデコード
	var req models.AssistantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Gemini APIクライアントの初期化
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Printf("Error creating client: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// モデルの設定
	model := client.GenerativeModel("gemini-2.0-flash")
	model.SetTemperature(0.7)
	model.SetMaxOutputTokens(2000)

	// レスポンスの生成
	resp, err := model.GenerateContent(ctx, genai.Text(req.Message))
	if err != nil {
		log.Printf("Error generating content: %v", err)
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成
	response := models.AssistantResponse{
		Response: string(resp.Candidates[0].Content.Parts[0].(genai.Text)),
	}

	// JSONレスポンスの送信
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
