package service

import (
	"GoLang_Tutorial/internal/dto"
	"GoLang_Tutorial/internal/models"
	"GoLang_Tutorial/internal/repository"
	"context"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetProduct(ctx context.Context, id uint) (*dto.ProductResponse, error)
}

type productServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}

func (s *productServiceImpl) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := &models.Product{
		Name:        req.Name,
		Code:        req.Code,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Code:  product.Code,
		Price: product.Price,
		Stock: product.Stock,
	}, nil
}

func (s *productServiceImpl) GetProduct(ctx context.Context, id uint) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &dto.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Code:  product.Code,
		Price: product.Price,
		Stock: product.Stock,
	}

	if product.Category != nil {
		res.CategoryName = product.Category.Name
	}

	return res, nil
}
