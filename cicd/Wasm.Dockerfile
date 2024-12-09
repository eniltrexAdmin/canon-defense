# the reason for this is to have always the same builder.
FROM golang:1.22 AS builder

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    gcc \
    libc6-dev \
    libgl1-mesa-dev \
    libxcursor-dev \
    libxi-dev \
    libxinerama-dev \
    libxrandr-dev \
    libxxf86vm-dev \
    libasound2-dev \
    pkg-config


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN GOOS=js GOARCH=wasm go build -o canon-defense ./ebiten