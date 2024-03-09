package repository

import (
	"context"
	paymentmethod "go-moduler/internal/payment_method"
)

type PaymentMethodRepository interface {
	GetAll(ctx context.Context, journey string) (paymentmethod.PaymentMethods, error)
	FindById(ctx context.Context, id uint64) (paymentmethod.PaymentMethod, error)
}
