package handlers

import (
	"io"
	"net/http"
)

// HandleHome はホームエンドポイントのハンドラーです
func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the API Server"}`))
}

// HandleHealth はヘルスチェックエンドポイントのハンドラーです
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "healthy"}`))
}

// HandleEcho はエコーエンドポイントのハンドラーです
func HandleEcho(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディを読み取る
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")

	// リクエストボディをそのまま返す
	w.Write(body)
}
