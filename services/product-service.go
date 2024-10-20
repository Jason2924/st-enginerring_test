package services

import (
	"context"
	"fmt"

	ntt "github.com/Jason2924/st-enginerring_test/entities"
	"github.com/Jason2924/st-enginerring_test/graph/model"
	rep "github.com/Jason2924/st-enginerring_test/repositories"
)

type ProductService interface {
	InsertFromFile(reqts []*ntt.ProductSchema) error
	ListMany(ctxt context.Context, reqt *ntt.ProductListManyReqt) ([]*model.Product, error)
}

type productService struct {
	productRepository rep.ProductRepository
}

func NewProductService(pdtRepo rep.ProductRepository) ProductService {
	return &productService{
		productRepository: pdtRepo,
	}
}

func (svc *productService) InsertFromFile(reqts []*ntt.ProductSchema) error {
	erro := svc.productRepository.InsertFromFile(reqts)
	return erro
}

func (svc *productService) ListMany(ctxt context.Context, reqt *ntt.ProductListManyReqt) ([]*model.Product, error) {
	currs := map[string]string{
		"USD": "$",
	}
	pdts, erro := svc.productRepository.ListMany(ctxt, reqt)
	if erro != nil {
		return nil, erro
	}
	resp := make([]*model.Product, 0, len(pdts))
	for _, item := range pdts {
		temp := &model.Product{
			ID:    item.ID,
			Name:  item.Name,
			Price: fmt.Sprintf("%s%f", currs[item.Currency], item.Price),
			Image: item.Image,
			Rating: &model.Rating{
				Average: item.RatingAverage,
				Reviews: item.RatingReviews,
			},
		}
		resp = append(resp, temp)
	}
	return resp, nil
}
