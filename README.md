# Chat Server

シンプルなGolang APIサーバー

## プロジェクト構造

```
.
├── server/     # サーバーコード
│   └── main.go
└── client/     # クライアントコード
    └── main.go
```

## 必要条件

- Go 1.21以上

## 実行方法

### サーバーの起動

```bash
cd server
go run main.go
```

サーバーは http://localhost:8080 で起動します。

### クライアントの実行

別のターミナルで以下のコマンドを実行します：

```bash
cd client
go run main.go
```

## エンドポイント

- `GET /`: ウェルカムメッセージ
- `GET /api/health`: ヘルスチェック
- `POST /api/echo`: リクエストボディをそのままエコーバック

## 開発

このプロジェクトは基本的なAPIサーバーのスケルトンとして機能します。必要に応じて新しいエンドポイントや機能を追加できます。

### クライアントの使用方法

クライアントは`client/main.go`に実装されており、以下の機能を提供します：

- `GetHome()`: ホームエンドポイントへのGETリクエスト
- `GetHealth()`: ヘルスチェックエンドポイントへのGETリクエスト
- `PostEcho(data interface{})`: エコーエンドポイントへのPOSTリクエスト

クライアントは`main()`関数で各エンドポイントのテストを実行します。 