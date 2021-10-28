package grpc_server

import (
	"context"
	"fmt"
	"net"
	"strconv"

	container "github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	pb "github.com/hramov/jobhelper/src/modules/grpc/proto"
	"github.com/hramov/jobhelper/src/modules/logger"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedJobhelperServer
}

func (s *Server) GetAllDevices(ctx context.Context, in *pb.GetAllRequest) ([]*pb.DeviceReply, error) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	var pbDevices []*pb.DeviceReply

	devices, err := deviceEntity.ShowAllDevices()

	if err != nil {
		logger.Log("GRPC", err.Error())
	}

	for _, device := range devices {
		id := strconv.Itoa(int(device.ID))
		prev_check := fmt.Sprintf("%v", device.PrevCheck)
		next_check := fmt.Sprintf("%v", device.NextCheck)

		pbDevices = append(pbDevices, &pb.DeviceReply{
			Id:          &id,
			Type:        &device.Type,
			Title:       &device.Title,
			Description: &device.Description,
			InvNumber:   &device.InvNumber,
			Location:    &device.Location,
			Status:      &device.Status,
			PrevCheck:   &prev_check,
			NextCheck:   &next_check,
			TagImageUrl: &device.TagImageUrl,
		})
	}

	return pbDevices, nil
}

func (s *Server) Start() {
	logger.Log("GRPC", "Server has beed started")
	lis, err := net.Listen("tcp", ":5005")
	if err != nil {
		logger.Log("GRPC", fmt.Sprintf("Failed to listen: %v", err))
	}
	server := grpc.NewServer()
	pb.RegisterJobhelperServer(server, &Server{})
	if err := server.Serve(lis); err != nil {
		logger.Log("GRPC", fmt.Sprintf("Failed to serve: %v", err))
	}
}
