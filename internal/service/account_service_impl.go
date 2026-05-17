package service

import (
	"context"
	_ "errors"
	"fmt"
	"time"

	"GoLang_Tutorial/internal/config"
	"GoLang_Tutorial/internal/dto"
	"GoLang_Tutorial/internal/models"
	"GoLang_Tutorial/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type accountServiceImpl struct {
	repo repository.AccountRepository
	cfg  *config.Config
}

func NewAccountService(repo repository.AccountRepository, cfg *config.Config) AccountService {
	return &accountServiceImpl{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *accountServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := &models.Account{
		ID:        uuid.New(),
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	// Bắt đầu Transaction
	err = s.repo.(*repository.AccountRepositoryPostgres).GetDB().Gorm.Transaction(func(tx *gorm.DB) error {
		// Bước 1: Lưu Account
		if err := tx.Create(account).Error; err != nil {
			return err
		}

		// Bước 2: Tạo Profile mặc định gắn với Account này
		user := &models.User{
			ID:        uuid.New(),
			AccountID: account.ID,
			FullName:  "", // Đợi người dùng cập nhật sau
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("không thể đăng ký: %w", err)
	}

	return &dto.RegisterResponse{
		ID:        account.ID,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
	}, nil
}

func (s *accountServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	account, err := s.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, models.ErrAccountNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return nil, models.ErrInvalidPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id": account.ID,
		"exp":        time.Now().Add(s.cfg.JWT.TokenExpiry).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: tokenString,
	}, nil
}
