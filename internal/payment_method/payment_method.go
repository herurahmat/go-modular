package paymentmethod

import (
	"errors"
	paymentcategory "go-moduler/internal/payment_category"
	pkgVo "go-moduler/pkg/valueobject"
)

type PaymentMethods []*PaymentMethod

type PaymentMethod struct {
	Id                uint64
	PaymentCategory   paymentcategory.PaymentCategory
	Name              string
	Code              string
	ExpiryDuration    pkgVo.ExpiryDuration
	PaymentMethodInfo PaymentMethodInfo
	MinimumAmount     float64
}

func (p PaymentMethod) IsEmpty() bool {
	return p.Id == 0 && p.Name == "" && p.Code == ""
}

func (p PaymentMethods) FilterMinimumAmount(amount float64) (*PaymentMethods, error) {

	paymentMethods := PaymentMethods{}
	for _, i := range p {
		if i.MinimumAmount <= amount {
			continue
		}
		paymentMethods = append(paymentMethods, i)
	}

	if len(paymentMethods) == 0 {
		return nil, errors.New("no payment method found")
	}

	return &paymentMethods, nil
}

func (p PaymentMethod) NewPaymentMethod() (*PaymentMethod, error) {

	paymentmethod := &PaymentMethod{
		Id:                p.Id,
		PaymentCategory:   p.PaymentCategory,
		Name:              p.Name,
		Code:              p.Code,
		ExpiryDuration:    p.ExpiryDuration,
		PaymentMethodInfo: p.PaymentMethodInfo,
		MinimumAmount:     p.MinimumAmount,
	}

	if paymentmethod.IsEmpty() {
		return paymentmethod, errors.New("PaymentMethod is empty")
	}

	return paymentmethod, nil

}
