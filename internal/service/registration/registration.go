package registration

import (
	"context"
	"database/sql"

	configs "elotus/config"
	"elotus/internal/model"
	db "elotus/internal/repo/dbmodel"
	convert_type "elotus/pkg/util/convert-type"
)

type (
	Service struct {
		conf *configs.Config

		DatabaseConn *sql.DB
		Query        *db.Queries

		jwtSvc jwtService
	}

	jwtService interface {
		GenerateTokenPair(userID int, userName string) (*model.TokenPair, error)
	}
)

func NewRegistrationService(conf *configs.Config, DatabaseConn *sql.DB, jwtSvc jwtService) *Service {
	query := db.New(DatabaseConn)
	return &Service{
		conf:         conf,
		DatabaseConn: DatabaseConn,
		Query:        query,
		jwtSvc:       jwtSvc,
	}
}

func (s *Service) CreateRegistration(ctx context.Context, req *model.CreateRegistrationRequest) (*model.CreateRegistrationResponse, error) {

	tx, err := s.DatabaseConn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	newUser := db.CreateUserParams{
		UserName: convert_type.NewNullString(req.UserName),
		Password: convert_type.NewNullString(req.Password),
	}
	err = s.Query.WithTx(tx).CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	user, err := s.Query.WithTx(tx).GetUser(ctx, convert_type.NewNullString(req.UserName))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tokenPair, err := s.jwtSvc.GenerateTokenPair(int(user.ID), user.UserName.String)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &model.CreateRegistrationResponse{
		TokenPair: *tokenPair,
	}, nil
}
