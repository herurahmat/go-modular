package repository

import (
	"context"
	"errors"
	paymentmethod "go-moduler/internal/payment_method"
)

type repositoryMysql struct {
	table string
}

// FindById implements PaymentMethodRepository.
func (r *repositoryMysql) FindById(ctx context.Context, id uint64) (paymentmethod.PaymentMethod, error) {
	panic("unimplemented")
}

// GetAll implements PaymentMethodRepository.
func (r *repositoryMysql) GetAll(ctx context.Context, journey string) (paymentmethod.PaymentMethods, error) {
	panic("unimplemented")
}

func NewRepositoryMysql(table string) (PaymentMethodRepository, error) {

	if table == "" {
		return nil, errors.New("no table payment method")
	}

	return &repositoryMysql{
		table: table,
	}, nil
}
