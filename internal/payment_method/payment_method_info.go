package paymentmethod

import "errors"

type PaymentMethodInfo struct {
	bankName                  string
	bankCode                  string
	paymentName               string
	merchantProfileFundSource string
	paymentMethodName         string
}

func (p PaymentMethodInfo) validation() error {

	if p.bankName == "" {
		return errors.New("bankName cannot be empty")
	}

	if p.bankCode == "" {
		return errors.New("bankCode cannot be empty")
	}

	if p.paymentName == "" {
		return errors.New("paymentName cannot be empty")
	}

	if p.merchantProfileFundSource == "" {
		return errors.New("merchantProfileFundSource cannot be empty")
	}

	return nil

}

func (p PaymentMethodInfo) IsEmpty() bool {
	return p.validation() != nil
}

func (p PaymentMethodInfo) BankName() string {
	return p.bankName
}

func (p PaymentMethodInfo) BankCode() string {
	return p.bankCode
}

func (p PaymentMethodInfo) PaymentName() string {
	return p.paymentName
}

func (p PaymentMethodInfo) FundSource() string {
	return p.merchantProfileFundSource
}

func (p PaymentMethodInfo) PaymentMethod() string {
	return p.paymentMethodName
}

func NewPaymentMethodInfo(
	bankName string,
	bankCode string,
	paymentName string,
	merchantProfileFundSource string,
	paymentMethodName string,
) (PaymentMethodInfo, error) {

	paymentMethodInfo := PaymentMethodInfo{
		bankName:                  bankName,
		bankCode:                  bankCode,
		paymentName:               paymentName,
		merchantProfileFundSource: merchantProfileFundSource,
		paymentMethodName:         paymentMethodName,
	}

	if err := paymentMethodInfo.validation(); err != nil {
		return PaymentMethodInfo{}, err
	}

	return paymentMethodInfo, nil

}
