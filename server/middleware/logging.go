package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware はリクエストのログを出力するミドルウェアです
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// リクエスト情報のログ出力
		log.Printf("Started %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		// 次のハンドラを実行
		next.ServeHTTP(w, r)

		// レスポンス時間の計算とログ出力
		duration := time.Since(start)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	})
}
