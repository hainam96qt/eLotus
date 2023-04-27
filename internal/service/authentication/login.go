package authentication

import (
	"context"
	"errors"

	"elotus/internal/model"
	convert_type "elotus/pkg/util/convert-type"
	"elotus/pkg/util/password"
)

func (s *Service) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.Query.GetUser(ctx, convert_type.NewNullString(req.UserName))
	if err != nil {
		return nil, err
	}

	ok := password.CheckPasswordHash(user.Password.String, req.Password)
	if !ok {
		return nil, errors.New("invalid password")
	}

	tokenPair, err := s.jwtSvc.GenerateTokenPair(int(user.ID), user.UserName.String)
	if err != nil {

		return nil, err
	}

	return &model.LoginResponse{
		TokenPair: *tokenPair,
	}, nil
}
