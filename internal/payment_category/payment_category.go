package paymentcategory

import "errors"

type PaymentCategories []*PaymentCategory

type PaymentCategory struct {
	Id   uint64
	Name string
}

func (p PaymentCategory) IsEmpty() bool {
	return p.Id == 0 && p.Name == ""
}

func NewPaymentCategory(id uint64, name string) (*PaymentCategory, error) {

	paymentCategory := PaymentCategory{Id: id, Name: name}

	if paymentCategory.IsEmpty() {
		return &paymentCategory, errors.New("PaymentCategory is empty")
	}

	return &paymentCategory, nil

}
