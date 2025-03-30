package models

// AssistantRequest はクライアントからのリクエストの構造体です
type AssistantRequest struct {
	Message string `json:"message"`
}

// AssistantResponse はアシスタントからのレスポンスの構造体です
type AssistantResponse struct {
	Response string `json:"response"`
}
