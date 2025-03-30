package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// ルーティングの設定
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/health", handleHealth)
	http.HandleFunc("/api/echo", handleEcho)

	// サーバーの起動
	log.Println("サーバーを起動します: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the API Server"}`))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "healthy"}`))
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
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
