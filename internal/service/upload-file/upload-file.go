package upload_file

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"

	configs "elotus/config"
)

// Size constants
const (
	maxSizeUpload = 8 * 1 << 20
)

type (
	Service struct {
		conf *configs.Config
	}
)

func NewUploadFileService(conf *configs.Config) *Service {
	return &Service{
		conf: conf,
	}
}

func (s *Service) UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > maxSizeUpload {
		return errors.New("images larger than 8 mb")
	}
	tempFile, err := os.CreateTemp("./temp", "upload-*.png")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	tempFile.Write(fileBytes)

	return nil
}
