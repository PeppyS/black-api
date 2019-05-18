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

func (bd *BusinessDelivery) Get(ctx context.Context, req *proto.GetBusinessRequest) (*proto.Business, error) {
	business, err := bd.service.GetByID(req.Id)
	if err != nil {
		return &proto.Business{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &proto.Business{
		Id:             business.ID,
		DisplayName:    business.DisplayName,
		DisplayAddress: business.DisplayAddress,
		AddressDetails: &proto.Business_AddressDetails{
			AddressLine_1: business.AddressLine1,
			AddressLine_2: business.AddressLine2.String,
			City:          business.AddressCity,
			State:         business.AddressState,
			PostalCode:    business.AddressPostalCode,
		},
	}, nil
}

func (bd *BusinessDelivery) List(ctx context.Context, req *proto.ListBusinessesRequest) (*proto.ListBusinessesResponse, error) {
	businesses, err := bd.service.List()
	if err != nil {
		return &proto.ListBusinessesResponse{}, status.Errorf(codes.Internal, "There was an unexpected internal server error: %s", err.Error())
	}

	response := &proto.ListBusinessesResponse{
		Businesses: []*proto.Business{},
	}

	for _, business := range businesses {
		response.Businesses = append(response.Businesses, &proto.Business{
			Id:             business.ID,
			DisplayName:    business.DisplayName,
			DisplayAddress: business.DisplayAddress,
			AddressDetails: &proto.Business_AddressDetails{
				AddressLine_1: business.AddressLine1,
				AddressLine_2: business.AddressLine2.String,
				City:          business.AddressCity,
				State:         business.AddressState,
				PostalCode:    business.AddressPostalCode,
			},
		})
	}

	return response, nil
}
