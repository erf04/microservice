package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"product/schema"
	"product/service"
)

type ProductHandler struct {
	service service.ProductService
}

func New(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (p *ProductHandler) GetProductByID(req []byte) (res []byte, err error) {
	fmt.Println("in the product handler")
	var body schema.GetProductByIDSchema
	if err := json.Unmarshal(req, &body); err != nil {
		return nil, err
	}
	ctx := context.Background()
	product, err := p.service.GetProductByID(ctx, body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(product)
}


func (p *ProductHandler) GetProducts(req []byte) (res []byte,err error) {
	var body schema.GetProductsSchema
	if err:= json.Unmarshal(req,&body);err != nil {
		return nil,err
	}
	ctx := context.Background()
	products , err := p.service.GetProducts(ctx,body)
	if err!= nil {
		return nil,err
	}
	return json.Marshal(products)
}