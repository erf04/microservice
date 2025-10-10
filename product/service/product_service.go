package service

import (
	"context"
	"product/models"
	"product/repository"
)

type ProductService interface {
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	GetProducts(ctx context.Context, body models.ProductWriteBody) ([]*models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *productService) GetProducts(ctx context.Context, body models.ProductWriteBody) ([]*models.Product, error) {
	return s.repo.Find(ctx, body)
}
