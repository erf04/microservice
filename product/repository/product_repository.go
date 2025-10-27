package repository

import (
	"context"
	"fmt"
	"time"

	"product/logger"
	"product/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"product/schema"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	FindByID(ctx context.Context, id string) (*models.Product, error)
	Find(ctx context.Context, body schema.GetProductsSchema) ([]*models.Product, error)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &productRepository{
		collection: db.Collection("products"),
	}
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Convert string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(id);
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Find(ctx context.Context, body schema.GetProductsSchema) ([]*models.Product, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{}
	nameProvided := body.Name != ""
	priceProvided := body.Price != 0

	if nameProvided {
		filter["name"] = bson.M{"$regex": body.Name, "$options": "i"}
	}
	if priceProvided {
		filter["price"] = body.Price
	}

	// If neither field is provided, return all products
	products, err := r.collection.Find(ctx, filter)
	if err != nil {
		logger.Logger.Fatalln("Error finding products:", err)
		return nil, err
	}
	// logger.Logger.Println("products count" , products.)
	var productsList []*models.Product
	for products.Next(ctx) {
		var p models.Product
		if err := products.Decode(&p); err != nil {
			return nil, err
		}
		productsList = append(productsList, &p)
	}
	logger.Logger.Println("product list count" , len(productsList))
	return productsList, nil
}
