# FROM golang:1.21-alpine

# # ディレクトリ指定
# ENV ROOT=/go/src/app
# WORKDIR ${ROOT}

FROM golang:1.23-alpine

# 必要なツールのインストール
RUN apk update && apk add --no-cache git curl

# 作業ディレクトリの指定
ENV ROOT=/go/src/app
WORKDIR ${ROOT}

# Goモジュールの有効化
ENV GO111MODULE=on