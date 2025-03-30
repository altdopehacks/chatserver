package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"chatserver/server/handlers"
)

func init() {
	// .envファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .envファイルが読み込めません: %v", err)
	}
}

func main() {
	// ルーティングの設定
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/api/health", handlers.HandleHealth)
	http.HandleFunc("/api/echo", handlers.HandleEcho)
	http.HandleFunc("/api/assistant", handlers.HandleAssistant)

	// サーバーの起動
	log.Println("サーバーを起動します: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
