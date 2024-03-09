package main

import (
	"fmt"
	filteroffer "go-moduler/internal/filter_offer"
	paymentMethodHandler "go-moduler/internal/payment_method/handler"
	"go-moduler/internal/payment_method/repository"
	paymentMethodService "go-moduler/internal/payment_method/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type infrastructures struct {
	filterOffer filteroffer.FilterOffer
}

type repositories struct {
	paymentMethodRepository repository.PaymentMethodRepository
}

type services struct {
	paymentMethodService paymentMethodService.PaymentMethodService
}

type domainHandlers struct {
	paymentMethodHandler *paymentMethodHandler.Handler
}

func newInfrastructures() (infrastructures, error) {

	filterOffer, err := filteroffer.NewFilterOfferHttp("EXAMPLE")

	if err != nil {
		return infrastructures{}, err
	}

	return infrastructures{
		filterOffer: filterOffer,
	}, nil
}

func newRepository() (repositories, error) {

	paymentMethodRepository, err := repository.NewRepositoryMysql(os.Getenv("TABLE"))

	if err != nil {
		return repositories{}, err
	}

	return repositories{
		paymentMethodRepository: paymentMethodRepository,
	}, nil
}

func newService(r repositories, i infrastructures) services {

	paymentMethodService := paymentMethodService.NewPaymentMethodService(r.paymentMethodRepository, i.filterOffer)

	return services{
		paymentMethodService: paymentMethodService,
	}
}

func router(h domainHandlers) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/payment_method", h.paymentMethodHandler.ModemHandler).Methods(http.MethodGet)

	return r
}

func main() {

	repository, err := newRepository()

	if err != nil {
		log.Panicln(err)
	}

	infrastructures, err := newInfrastructures()

	if err != nil {
		log.Panicln(err)
	}

	services := newService(repository, infrastructures)

	(&http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		Handler: handlers.CompressHandler(
			router(domainHandlers{
				paymentMethodHandler.NewHandler(services.paymentMethodService),
			}),
		),
		ReadHeaderTimeout: 3 * time.Second,
	}).ListenAndServe()

}
