﻿# ProtoBuf-Chat-Bidirecional

Este é um projeto de chat em tempo real com **gRPC** usando **stream bidirecional**, desenvolvido em Go.

## 📦 Estrutura do Projeto

## 1. Clonar o repositório

```bash
git clone git@github.com:ZecaSouza/ProtoBuf-Chat-Bidirecional.git
cd ProtoBuf-Chat-Bidirecional
```

## 2. Instalar dependências
Go (versão recomendada: 1.18+)

```bash
go mod tidy
```
## 3. Compilar o .proto
```bash
# Instalar o compilador Protocol Buffers
brew install protobuf

# Plugin do Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Adicione o diretório do Go bin ao PATH se ainda não estiver:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
Compile o proto:
```bash
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative chat/chat.proto
```

## 4. Rodar o servidor
```bash
go run server.go
```

## 5. Rodar um cliente (em outro terminal)
```bash
go run client.go
```
