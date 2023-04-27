package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	configs "elotus/config"
	"elotus/internal/model"
)

type Service struct {
	*configs.Config
}

type Claims struct {
	Username string `json:"username"`
	USerID   int    `json:"user_id"`
	jwt.StandardClaims
}

func NewJwtService(config *configs.Config) *Service {
	return &Service{
		config,
	}
}

func (s *Service) GenerateTokenPair(userID int, userName string) (*model.TokenPair, error) {
	accessToken, accessTokenExpiresAt, err := s.generateToken(userID, userName, s.Config.Jwt.AccessTokenTLL)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshTokenExpiresAt, err := s.generateToken(userID, userName, s.Config.Jwt.AccessTokenTLL)
	if err != nil {
		return nil, err
	}
	return &model.TokenPair{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
	}, nil
}

func (s *Service) generateToken(userID int, userName string, duration time.Duration) (string, time.Time, error) {
	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		Username: userName,
		USerID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.Jwt.SecretKey))
	if err != nil {
		return "", expirationTime, err
	}
	return tokenString, expirationTime, nil
}
