package filteroffer

import "context"

type FilterOffer interface {
	FindSubMenuLevel(ctx context.Context, subMenuLevel string) (FilterOfferData, error)
}

type FilterOfferData struct {
	Id    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
