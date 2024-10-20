package repositories

import (
	"context"

	dtb "github.com/Jason2924/st-enginerring_test/databases"
	ntt "github.com/Jason2924/st-enginerring_test/entities"
)

type ProductRepository interface {
	InsertFromFile(reqts []*ntt.ProductSchema) error
	ListMany(ctxt context.Context, reqt *ntt.ProductListManyReqt) ([]*ntt.ProductSchema, error)
}

type productRepository struct {
	mysqlDatabase dtb.MysqlDatabase
}

func NewProductRepository(msqDtbs dtb.MysqlDatabase) ProductRepository {
	return &productRepository{
		mysqlDatabase: msqDtbs,
	}
}

func (rep *productRepository) InsertFromFile(reqts []*ntt.ProductSchema) error {
	conn := rep.mysqlDatabase.Connect()
	qery := conn.Model(ntt.ProductSchema{}).Create(reqts)
	return qery.Error
}

func (rep *productRepository) ListMany(ctxt context.Context, reqt *ntt.ProductListManyReqt) ([]*ntt.ProductSchema, error) {
	resp := []*ntt.ProductSchema{}
	conn := rep.mysqlDatabase.Connect().WithContext(ctxt)
	qery := conn.Model(ntt.ProductSchema{})

	if reqt.Limit > 0 && reqt.Page > 0 {
		qery.Offset(reqt.Limit * (reqt.Page - 1)).Limit(reqt.Limit)
	}
	rsul := qery.Find(&resp)
	if rsul.Error != nil {
		return nil, rsul.Error
	}
	return resp, nil
}
