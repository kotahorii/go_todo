
FROM golang:1.16

WORKDIR /app
# go mod init xxx でgo.modファイルを作成しておくこと
COPY go.mod .
COPY go.sum .

# go modからパッケージをダウンロード
RUN go mod download

# /app にすべてのコードをコピー
COPY . .

# エントリポイント
CMD ["go", "run", "main.go"]