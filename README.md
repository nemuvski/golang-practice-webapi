# golang-practice-webapi

`echo` と Redis(`go-redis`をクライアントに利用) を用いたレート制限付きの WebAPI サーバを作成するサンプルプログラム

## 必要なソフトウェア等

- Golang `v1.20`
- Docker

## 準備

```bash
# Run redis server
docker-compose up -d

# Create dotenv file
cp .env.example .env.local
```

```bash
go mod tidy

go run ./src
```

## WebAPI の実行例

```bash
# /v1/public/:id
# :idは0-255の整数値をとる
curl --location 'localhost:8080/v1/public/32'
```

```bash
# /v1/private/:id
# :idは0-255の整数値をとる
curl --location 'localhost:8080/v1/private/100' \
--header 'Authorization: Bearer a1b2c3d4e5f6g7h8i9j0k'
# Bearerの後の値は.env.local中のAPP_API_VALID_TOKENの内容を設定
```
