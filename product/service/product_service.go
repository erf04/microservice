package service

import (
	"context"
	"product/models"
	"product/repository"
	"product/schema"
)

type ProductService interface {
	GetProductByID(ctx context.Context, body schema.GetProductByIDSchema) (*models.Product, error)
	GetProducts(ctx context.Context, body schema.GetProductsSchema) ([]*models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProductByID(ctx context.Context, body schema.GetProductByIDSchema) (*models.Product, error) {
	return s.repo.FindByID(ctx, body.ID)
}

func (s *productService) GetProducts(ctx context.Context, body schema.GetProductsSchema) ([]*models.Product, error) {
	return s.repo.Find(ctx, body)
}
