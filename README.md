# cache-service

# Requirements

go protoc and docker will be required to run the program so make sure to to install them. can refer here
To install go https://go.dev/doc/install
To install protoc https://grpc.io/docs/protoc-installation/
To install docker https://docs.docker.com/engine/install/

# How to run the project

1. lets start with starting a redis server using docker. here are the steps  
  a. "docker build -t redis-cache . " to build docker image.
  b. "docker run -p '6379:6379' -d redis-cache" to run the run redis images.
2. now you start the server with "go run server/server.go" make sure to be in server directory
3. after that you can run client with "go run cllient/client.go" 

note: current client code calls each rpc method once. it calls Set user with following credentials 
    Name:     Alice
    class     IV
    RollNum   15
    MetaData  secret message alice is wonder full

note: make sure z_generated is populated with grpc auto generated code if not it can be done with running following command in terminal (or cmd) make sure to be in proto directory before running this command:

  $ protoc --go_out=../z_generated --go_opt=paths=source_relative \
           --go-grpc_out=../z_generated --go-grpc_opt=paths=source_relative \
              cache.proto
  
# Assumption while implementing this code 
1. The user is Unique with respect to Name and RollNo attributes. For eg. Name Alice and RollNo 15 are not present with two different class like IV and V since the requirement said that Getuser will be retrived based on Name and RollNo credentials. 

# if you want to create customer client then refer this 
1. your run the server code as mention in the How to run the project section.
2. to implement custom client you can use ConnectToCache function from client/client package.
ConnectTocache function return a CacheClient object if the connect to cache service is successfull or else it returns error 
if the connection is successfull. You can use CacheClient object to call rpc method based on requirements with respective arguments.
