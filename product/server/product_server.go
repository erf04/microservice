package server

import (
	"context"
	pb "product/proto/productpb"
	"product/service"
	"product/models"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	svc service.ProductService
}

func NewProductServer(svc service.ProductService) *ProductServer {
	return &ProductServer{svc: svc}
}

func (s *ProductServer) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.GetProductResponse, error) {
	product, err := s.svc.GetProductByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		},
	}, nil
}


func (s *ProductServer) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	body := models.ProductWriteBody{
		Name:  req.GetName(),
		Price: (float64)(req.GetPrice()),
	}
	products, err := s.svc.GetProducts(ctx, body)
	if err != nil {
		return nil, err
	}

	var productResponses []*pb.GetProductResponse
	for _, product := range products {
		productResponses = append(productResponses, &pb.GetProductResponse{
			Product: &pb.Product{
				Id:    product.ID,
				Name:  product.Name,
				Price: product.Price,
			},
		})
	}

	return &pb.GetProductsResponse{
		Products: productResponses,
	}, nil
}