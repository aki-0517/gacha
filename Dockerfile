# 使用するGoのバージョンを指定
FROM golang:1.18

# 作業ディレクトリを設定
WORKDIR /app

# ホストのファイルをコンテナにコピー
COPY . .

# 依存関係をインストール
RUN go mod download

# アプリケーションをビルド
RUN go build -o main .

# 実行コマンド
CMD ["./main"]
