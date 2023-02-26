package service

import (
	"context"

	pb "kratosent/api/mtest"
)

type V1Service struct {
	pb.UnimplementedV1Server
}

func NewV1Service() *V1Service {
	return &V1Service{}
}

func (s *V1Service) CreateV1(ctx context.Context, req *pb.CreateV1Request) (*pb.CreateV1Reply, error) {
	return &pb.CreateV1Reply{}, nil
}
func (s *V1Service) UpdateV1(ctx context.Context, req *pb.UpdateV1Request) (*pb.UpdateV1Reply, error) {
	return &pb.UpdateV1Reply{}, nil
}
func (s *V1Service) DeleteV1(ctx context.Context, req *pb.DeleteV1Request) (*pb.DeleteV1Reply, error) {
	return &pb.DeleteV1Reply{}, nil
}
func (s *V1Service) GetV1(ctx context.Context, req *pb.GetV1Request) (*pb.GetV1Reply, error) {
	return &pb.GetV1Reply{}, nil
}
func (s *V1Service) ListV1(ctx context.Context, req *pb.ListV1Request) (*pb.ListV1Reply, error) {
	return &pb.ListV1Reply{}, nil
}
