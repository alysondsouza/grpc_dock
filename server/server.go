package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"GO_GRPC_CHAT/chat"	//?
)

// Server:
// docker build -t myDocker --no-cache .
// docker run -p 9000:9000 -tid myDocker
// returns a connetionString: 56dc91ca552555a0f099defc265635d127f5692abb86b8432b7a61c9eb99aee9
// Client:
// Run go .

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChittyChatServiceServer(grpcServer, &s)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
