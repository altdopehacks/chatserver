package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"chatserver/server/handlers"
	"chatserver/server/middleware"
)

func init() {
	// .envファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .envファイルが読み込めません: %v", err)
	}
}

func main() {
	// ルーティングの設定
	mux := http.NewServeMux()

	// エンドポイントの設定
	mux.HandleFunc("/", handlers.HandleHome)
	mux.HandleFunc("/api/health", handlers.HandleHealth)
	mux.HandleFunc("/api/echo", handlers.HandleEcho)
	mux.HandleFunc("/api/assistant", handlers.HandleAssistant)

	// ミドルウェアの適用（内側から外側の順に実行される）
	handler := middleware.LoggingMiddleware(
		middleware.CORSMiddleware(mux),
	)

	// サーバーの起動
	log.Println("サーバーを起動します: http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
