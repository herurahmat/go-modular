package service

import (
	"context"
	filteroffer "go-moduler/internal/filter_offer"
	paymentmethod "go-moduler/internal/payment_method"
	"go-moduler/internal/payment_method/repository"
)

type paymentMethodService struct {
	repository  repository.PaymentMethodRepository
	filterOffer filteroffer.FilterOffer
}

// PurchaseData implements PaymentMethodService.
func (p *paymentMethodService) PurchaseModem(ctx context.Context) (paymentmethod.PaymentMethods, error) {

	data, err := p.repository.GetAll(ctx, "PURCHASE_MODEM")

	if err != nil {
		return nil, err
	}

	return data, nil

}

// PurchaseModem implements PaymentMethodService.
func (p *paymentMethodService) PurchaseData(ctx context.Context, subMenuLevel string) (paymentmethod.PaymentMethods, error) {

	browseOffer, err := p.filterOffer.FindSubMenuLevel(ctx, subMenuLevel)

	if err != nil {
		return nil, err
	}

	data, err := p.repository.GetAll(ctx, "PURCHASE_DATA")

	if err != nil {
		return nil, err
	}

	filterMinimumAmount, err := data.FilterMinimumAmount(browseOffer.Price)

	return *filterMinimumAmount, err

}

func NewPaymentMethodService(
	repository repository.PaymentMethodRepository,
	filterOffer filteroffer.FilterOffer,
) PaymentMethodService {
	return &paymentMethodService{
		repository:  repository,
		filterOffer: filterOffer,
	}
}
