package app

import (
	"context"
	"errors"

	"github.com/joseluis8906/go-code/idl/dhpspb"
	"github.com/joseluis8906/go-code/src/internal/dhps/domain"
	"github.com/joseluis8906/go-code/src/pkg/delivery"
)

type (
	DeliveryService struct {
		dhpspb.UnimplementedDeliveryServer

		AssistantRepository AssistantRepository
		CustomerRepository  CustomerRepository
	}

	AssistantRepository interface {
		Get(context.Context) *delivery.Assistant
	}

	CustomerRepository interface {
		Get(ctx context.Context, email string) delivery.Customer
	}
)

// NewDeliveryService returns a new instance of DeliveryService.
func NewDeliveryService(params Params) *DeliveryService {
	return &DeliveryService{
		AssistantRepository: params.AssistantRepository,
		CustomerRepository:  params.CustomerRepository,
	}
}

// CustomerAsksForAProduct returns a list of suggestions.
func (s *DeliveryService) CustomerAsksForAProduct(ctx context.Context, req *dhpspb.CustomerAsksForAProductReq) (*dhpspb.CustomerAsksForAProductRes, error) {
	theAssistant := s.AssistantRepository.Get(ctx)
	theCustomer := s.CustomerRepository.Get(ctx, req.GetCustomerEmail())

	if theAssistant == nil || theCustomer.IsZero() {
		return nil, errors.New("nil assistant or customer")
	}

	action := domain.NewCustomerAsksForAProduct(&theCustomer, req.GetProductName(), theAssistant)
	suggestions := action.Do()

	res := &dhpspb.CustomerAsksForAProductRes{}
	for _, product := range suggestions {
		res.Products = append(res.Products, &dhpspb.Product{Ref: product.Ref, Name: product.Name})
	}

	return res, nil
}
