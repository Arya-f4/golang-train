package usecase

import (
	"back-train/internal/domain"
	"back-train/internal/repository"
	"back-train/pkg/utils"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type authUsecase struct {
	userRepo           repository.UserRepository
	jwtSecret          string
	jwtExpirationHours time.Duration
}

func NewAuthUsecase(ur repository.UserRepository, secret string, exp time.Duration) AuthUsecase {
	return &authUsecase{
		userRepo:           ur,
		jwtSecret:          secret,
		jwtExpirationHours: exp,
	}
}

func (u *authUsecase) Register(ctx context.Context, email, password string) (*domain.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}

	// Default role is 'user'
	return u.userRepo.CreateUser(ctx, user, "user")
}

func (u *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"roles":   user.Roles,
		"exp":     time.Now().Add(u.jwtExpirationHours).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
