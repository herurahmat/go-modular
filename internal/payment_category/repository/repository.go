package repository

import (
	"context"
	paymentcategory "go-moduler/internal/payment_category"
)

type PaymentCategoryRepository interface {
	GetAll(ctx context.Context) (paymentcategory.PaymentCategories, error)
	FindById(ctx context.Context, id uint64) (paymentcategory.PaymentCategory, error)
}
