package filteroffer

import "context"

type filterOfferHttp struct {
	exampleConfig string
}

// FindSubMenuLevel implements FilterOffer.
func (f *filterOfferHttp) FindSubMenuLevel(ctx context.Context, subMenuLevel string) (FilterOfferData, error) {
	return FilterOfferData{
		Id:    1,
		Name:  "Name",
		Price: 1000,
	}, nil
}

func NewFilterOfferHttp(exampleConfig string) (FilterOffer, error) {

	return &filterOfferHttp{
		exampleConfig: exampleConfig,
	}, nil
}
