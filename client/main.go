package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "http://localhost:8080"

// Client はAPIクライアントの構造体です
type Client struct {
	httpClient *http.Client
}

// NewClient は新しいクライアントを作成します
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// GetHome はホームエンドポイントにGETリクエストを送信します
func (c *Client) GetHome() (string, error) {
	resp, err := c.httpClient.Get(baseURL + "/")
	if err != nil {
		return "", fmt.Errorf("リクエストエラー: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("レスポンスの読み取りエラー: %v", err)
	}

	return string(body), nil
}

// GetHealth はヘルスチェックエンドポイントにGETリクエストを送信します
func (c *Client) GetHealth() (string, error) {
	resp, err := c.httpClient.Get(baseURL + "/api/health")
	if err != nil {
		return "", fmt.Errorf("リクエストエラー: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("レスポンスの読み取りエラー: %v", err)
	}

	return string(body), nil
}

// PostEcho はエコーエンドポイントにPOSTリクエストを送信します
func (c *Client) PostEcho(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("JSONエンコードエラー: %v", err)
	}

	resp, err := c.httpClient.Post(
		baseURL+"/api/echo",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", fmt.Errorf("リクエストエラー: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("レスポンスの読み取りエラー: %v", err)
	}

	return string(body), nil
}

func main() {
	client := NewClient()

	// ホームエンドポイントのテスト
	fmt.Println("ホームエンドポイントのテスト:")
	homeResp, err := client.GetHome()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("レスポンス: %s\n", homeResp)
	}

	// ヘルスチェックのテスト
	fmt.Println("\nヘルスチェックのテスト:")
	healthResp, err := client.GetHealth()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("レスポンス: %s\n", healthResp)
	}

	// エコーエンドポイントのテスト
	fmt.Println("\nエコーエンドポイントのテスト:")
	testData := map[string]string{
		"message": "Hello, World!",
		"time":    "2024-03-30",
	}
	echoResp, err := client.PostEcho(testData)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
	} else {
		fmt.Printf("レスポンス: %s\n", echoResp)
	}
}
