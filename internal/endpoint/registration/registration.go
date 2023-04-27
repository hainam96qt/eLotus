package registration

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"elotus/internal/model"
	error2 "elotus/pkg/error"
	"elotus/pkg/util/request"
	"elotus/pkg/util/response"
)

type (
	Endpoint struct {
		registrationSvc registrationService
	}

	registrationService interface {
		CreateRegistration(ctx context.Context, req *model.CreateRegistrationRequest) (*model.CreateRegistrationResponse, error)
	}
)

func InitRegistrationHandler(r *chi.Mux, registrationSvc registrationService) {
	registrationEndpoint := &Endpoint{
		registrationSvc: registrationSvc,
	}
	r.Route("/api/v1/registration", func(r chi.Router) {
		r.Post("/", registrationEndpoint.createRegistrationRequest)
	})
}

func (e *Endpoint) createRegistrationRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateRegistrationRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))
		return
	}

	res, err := e.registrationSvc.CreateRegistration(ctx, &req)
	if err != nil {
		log.Printf("failed to register new user: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))
		return
	}

	response.JSON(w, res)
}
