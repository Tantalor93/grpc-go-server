package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"grpc-go-server/proto"
	"net"
)

type DnsForwardingDecisionService struct {
	proto.UnimplementedDNSForwardingDecisionServiceServer
}

func (s DnsForwardingDecisionService) EvaluateCustomDNSResolution(ctx context.Context, request *proto.EvaluationRequest) (*proto.EvaluationResponse, error) {
	resp := proto.EvaluationResponse{
		ZoneForwarding: true,
	}
	if len(request.CustomerId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid customer")
	}
	return &resp, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterDNSForwardingDecisionServiceServer(server, DnsForwardingDecisionService{})
	listen, _ := net.Listen("tcp", ":9111")
	reflection.Register(server)
	if err := server.Serve(listen); err != nil {
		fmt.Println(err)
	}
}

