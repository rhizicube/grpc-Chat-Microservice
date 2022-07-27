chat:
  export PATH="$PATH:$(go env GOPATH)/bin"
  protoc -Ichat/proto --go_out=. --go_opt=module=gRPC  --go-grpc_out=. --go-grpc_opt=module=gRPC  chat/proto/chat.proto
  go build -o bin/chat/server ./chat/server
  go build -o bin/chat/client ./chat/client


server:
  ./bin/chat/server


client:
  ./bin/chat/client
