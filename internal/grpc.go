package internal

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"schedule/pkg/pb"
)

type Server struct{}

func (s *Server) GetCommands(ctx context.Context, in *pb.Empty) (*pb.CommandsResp, error) {
	commands, err := GetCommands()
	if err != nil {
		return nil, err
	}
	PbCommands := make([]*pb.Command, 0)
	for _, cmd := range commands {
		SingleCommand := pb.Command{
			Id:         cmd.ID,
			Command:    cmd.Command,
			CreateTime: cmd.CreateTime.String(),
		}
		PbCommands = append(PbCommands, &SingleCommand)
	}
	return &pb.CommandsResp{Command: PbCommands}, nil
}

func (s *Server) GetScheduleOne(ctx context.Context, in *pb.GetScheduleOneReq) (*pb.GetScheduleOneResp, error) {
	resp, err := GetScheduleOne(in.Id)
	if err != nil {
		return nil, err
	}
	result := pb.GetScheduleOneResp{}
	result.Id = resp.ID
	result.TimeTypeId = resp.TimeTypeID
	result.IntervalDay = resp.IntervalDay
	result.IntervalSeconds = resp.IntervalSeconds
	result.AtTime = resp.AtTime
	result.StartTime = resp.StartTime
	result.EndTime = resp.EndTime
	result.CommandId = resp.CommandID
	result.Name = resp.Name
	result.StartDate = timestamppb.New(resp.StartDate)
	result.EndDate = timestamppb.New(resp.EndDate)
	result.Enable = resp.Enable
	result.Repeat = resp.Repeat
	result.CreateTime = timestamppb.New(resp.CreateTime)
	result.RepeatWeekday = resp.RepeatWeekday
	result.RepeatDay = resp.RepeatDay
	result.RepeatMonth = resp.RepeatMonth
	return &result, nil
}

func GrpcServer() {
	// Starts a TCP server listening on port 55555 and handles any errors.
	l, err := net.Listen("tcp", ":55555")
	// The gRPC server will use it.
	if err != nil {
		log.Fatalf("failed to listen for tcp: %s", err)
	}
	grpcServer := grpc.NewServer() // Creates a gRPC server and handles requests over the TCP connection
	pb.RegisterGetScheduleCRUDServiceServer(grpcServer, &Server{})
	grpcServer.Serve(l)
}
