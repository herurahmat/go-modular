package service

import (
	"context"
	paymentmethod "go-moduler/internal/payment_method"
)

type PaymentMethodService interface {
	PurchaseModem(ctx context.Context) (paymentmethod.PaymentMethods, error)
	PurchaseData(ctx context.Context, subMenuLevel string) (paymentmethod.PaymentMethods, error)
}
