package service

import (
	"context"
	"product-service/internal/model"
	"product-service/internal/repository"
)

type ProductService interface {
	Create(ctx context.Context, req model.CreateProductRequest) (*model.Product, error)
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	GetAll(ctx context.Context, limit, offset int) ([]model.Product, error)
	Update(ctx context.Context, id uint, req model.UpdateProductRequest) (*model.Product, error)
	Delete(ctx context.Context, id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(ctx context.Context, req model.CreateProductRequest) (*model.Product, error) {
	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *productService) GetAll(ctx context.Context, limit, offset int) ([]model.Product, error) {
	return s.repo.FindAll(ctx, limit, offset)
}

func (s *productService) Update(ctx context.Context, id uint, req model.UpdateProductRequest) (*model.Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}
	if req.Category != "" {
		product.Category = req.Category
	}
	if req.ImageURL != "" {
		product.ImageURL = req.ImageURL
	}

	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
