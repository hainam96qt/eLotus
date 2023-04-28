package upload_file

import (
	"context"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi/v5"

	"elotus/internal/model"
	error2 "elotus/pkg/error"
	"elotus/pkg/midleware"
	"elotus/pkg/util/response"
)

type (
	Endpoint struct {
		uploadFileSvc uploadFileService
	}

	uploadFileService interface {
		UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error
	}
)

func InitAuthenticationHandler(r *chi.Mux, uploadFileSvc uploadFileService) {
	endpoint := &Endpoint{
		uploadFileSvc: uploadFileSvc,
	}

	r.Route("/api/v1/files", func(r chi.Router) {
		r.Use(midleware.AuthenticateMW.Authenticator)
		r.Post("/upload", endpoint.upload)
	})
}

func (e *Endpoint) upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	identity, err := midleware.GetIdentityFromContext(ctx)
	if err != nil {
		log.Printf("read request body error: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))

	}
	_ = identity
	
	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("read request body error: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))

	}
	defer file.Close()

	err = e.uploadFileSvc.UploadFile(ctx, file, header)
	if err != nil {
		log.Printf("failed to upload file: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))
	}
	response.JSON(w, model.Success{Message: "successfully uploaded file"})
}
