package handler

import paymentmethod "go-moduler/internal/payment_method"

type ExpiryDuration struct {
	Amount float64
	Unit   string
}

type PaymentMethodResponse struct {
	Id              uint64 `json:"id"`
	PaymentCategory string `json:"payment_category"`
	Name            string `json:"name"`
	Code            string `json:"code"`
	ExpiryDuration  ExpiryDuration
	MinimumAmount   float64
}

type PaymentMethodResponses []PaymentMethodResponse

func responsePaymentMethod(data paymentmethod.PaymentMethods) PaymentMethodResponses {

	if data == nil {
		return PaymentMethodResponses{}
	}

	responses := PaymentMethodResponses{}

	for _, i := range data {
		responses = append(responses, PaymentMethodResponse{
			Id:              i.Id,
			PaymentCategory: i.PaymentCategory.Name,
			Name:            i.Name,
			Code:            i.Code,
			ExpiryDuration: ExpiryDuration{
				Amount: i.ExpiryDuration.Value(),
				Unit:   i.ExpiryDuration.UnitOrMeasurement(),
			},
			MinimumAmount: i.MinimumAmount,
		})

	}

	return responses
}
