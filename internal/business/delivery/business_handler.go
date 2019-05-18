package delivery

import (
	"context"

	"github.com/PeppyS/black-api/internal/business/service"
	"github.com/PeppyS/black-api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BusinessDelivery struct {
	service *service.BusinessService
}

func NewBusinessDelivery(bs *service.BusinessService) *BusinessDelivery {
	return &BusinessDelivery{bs}
}

func (bd *BusinessDelivery) List(ctx context.Context, req *proto.ListBusinessesRequest) (*proto.ListBusinessesResponse, error) {
	businesses, err := bd.service.List()
	if err != nil {

		return &proto.ListBusinessesResponse{}, status.Errorf(codes.Internal, "There was an unexpected internal server error: %s", err.Error())
	}

	response := &proto.ListBusinessesResponse{
		Businesses: []*proto.Business{},
	}

	for _, value := range businesses {
		response.Businesses = append(response.Businesses, &proto.Business{
			Id:             value.ID,
			DisplayName:    value.DisplayName,
			DisplayAddress: value.DisplayAddress,
			AddressDetails: &proto.Business_AddressDetails{
				AddressLine_1: value.AddressLine1,
				AddressLine_2: value.AddressLine2.String,
				City:          value.AddressCity,
				State:         value.AddressState,
				PostalCode:    value.AddressPostalCode,
			},
		})
	}

	return response, nil
}
