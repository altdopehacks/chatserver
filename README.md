# Chat Server and test client

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

## API仕様

### エンドポイント一覧

| メソッド | エンドポイント | 説明 |
|----------|----------------|------|
| GET | `/` | ウェルカムメッセージ |
| GET | `/api/health` | ヘルスチェック |
| POST | `/api/echo` | リクエストボディをそのままエコーバック |
| POST | `/api/assistant` | Gemini AIを使用した応答生成 |

### 各エンドポイントの詳細

#### GET /
ウェルカムメッセージを返します。

**レスポンス**
```json
{
    "message": "Welcome to the API Server"
}
```

#### GET /api/health
サーバーの健康状態を返します。

**レスポンス**
```json
{
    "status": "healthy"
}
```

#### POST /api/echo
送信されたリクエストボディをそのまま返します。

**リクエスト**
```json
{
    "message": "Hello, World!"
}
```

**レスポンス**
```json
{
    "message": "Hello, World!"
}
```

#### POST /api/assistant
Gemini AIを使用して応答を生成します。

**リクエスト**
```json
{
    "message": "Goプログラミングについて教えてください"
}
```

**レスポンス**
```json
{
    "response": "生成された応答テキスト"
}
```

### JavaScriptからの呼び出し例

```javascript
// アシスタントエンドポイントの呼び出し
async function callAssistant(message) {
    try {
        const response = await fetch('http://localhost:8080/api/assistant', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                message: message
            })
        });
        
        const data = await response.json();
        return data.response;
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
}

// 使用例
callAssistant('Goプログラミングについて教えてください')
    .then(response => console.log(response))
    .catch(error => console.error(error));
```

## 開発

このプロジェクトは基本的なAPIサーバーのスケルトンとして機能します。必要に応じて新しいエンドポイントや機能を追加できます。

### クライアントの使用方法

クライアントは`client/main.go`に実装されており、以下の機能を提供します：

- `GetHome()`: ホームエンドポイントへのGETリクエスト
- `GetHealth()`: ヘルスチェックエンドポイントへのGETリクエスト
- `PostEcho(data interface{})`: エコーエンドポイントへのPOSTリクエスト
- `PostAssistant(message string)`: アシスタントエンドポイントへのPOSTリクエスト

クライアントは`main()`関数で各エンドポイントのテストを実行します。 