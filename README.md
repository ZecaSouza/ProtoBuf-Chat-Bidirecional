# ProtoBuf-Chat-Bidirecional

Este é um projeto de chat em tempo real com **gRPC** usando **stream bidirecional**, desenvolvido em Go.

Ele utiliza o **Protocol Buffers (Protobuf)**, um formato de serialização de dados leve e eficiente criado pelo Google. O Protobuf permite definir a estrutura das mensagens e dos serviços em um arquivo `.proto`, gerando automaticamente o código necessário para comunicação entre cliente e servidor de forma rápida (em média até 5x mais veloz que JSON), compacta e multiplataforma.

As mensagens são convertidas para **formato binário**, o que reduz o tamanho dos dados transmitidos e acelera a comunicação. Diferente de formatos como JSON ou XML, que são baseados em texto e mais verbosos, o Protobuf gera arquivos menores e mais rápidos de serializar e desserializar.

Além disso, o **gRPC** utiliza o **HTTP/2** como protocolo de transporte, o que traz várias vantagens em relação ao HTTP/1.1, como:

- Multiplexação de streams (várias mensagens em uma mesma conexão)
- Redução de latência
- Cabeçalhos comprimidos
- Conexões persistentes mais eficientes

Com isso, o projeto consegue estabelecer uma comunicação contínua e bidirecional entre clientes e servidor de forma moderna, eficiente e escalável.



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


📌 Melhorias futuras
✅ Suporte a TLS (criptografia)

✅ Armazenamento de histórico (arquivo ou banco de dados)

✅ Interface gráfica com BubbleTea ou WebSocket

✅ Autenticação de usuários

✅ Argumentos CLI para configuração
