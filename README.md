# grpc-Chat-Microservice
A simple micro-service in which client and server chat whit each other using gRPC-go 
## Steps how to run this file
Their is a few step to run your chat-microservice
#### STEP 1:
  Download it and save it in gRPC dirctory and open your terminal
  
#### STEP 2   To Regenrate a gRPC code :
    write on terminal
   <pre>protoc -Ichat/proto --go_out=. --go_opt=module=gRPC  --go-grpc_out=. --go-grpc_opt=module=gRPC  chat/proto/chat.proto</pre>
   
#### STEP 3 To Built a server :
  <pre>go build -o bin/chat/server ./chat/server</pre>
  
#### STEP 4 To Built a client :
  <pre>go build -o bin/chat/client ./chat/client</pre>
  
#### STEP 5 To Run the Server :
  write on server termial
  <pre>./bin/chat/server</pre>
  
#### STEP 6 To Run the Client :
  write on client terminal
  <pre>./bin/chat/client</pre>
