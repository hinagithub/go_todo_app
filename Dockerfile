
# リリース用のビルドを行うコンテナイメージを作成するステージ
FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# --------------------------------------------------------

# デプロイ用のコンテナ
# マネージドサービス上で動かすことを想定したリリース用のコンテナイメージ を作成するステージ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# --------------------------------------------------------

# ローカル開発環境で利用するホットリロード環境
# ローカルで開発するときに利用するコンテナイメージを作成するステージ
FROM golang:1.18.2 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "80"]