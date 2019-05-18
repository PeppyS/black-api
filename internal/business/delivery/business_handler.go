package delivery

import (
	"context"

	"github.com/PeppyS/black-api/proto"
)

type BusinessDelivery struct {
}

func NewBusinessDelivery() *BusinessDelivery {
	return &BusinessDelivery{}
}

func (bd *BusinessDelivery) List(ctx context.Context, req *proto.ListBusinessesRequest) (*proto.ListBusinessesResponse, error) {
	return &proto.ListBusinessesResponse{}, nil
}
