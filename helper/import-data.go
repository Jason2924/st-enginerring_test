package helper

import (
	"encoding/csv"
	"os"

	ntt "github.com/Jason2924/st-enginerring_test/entities"
	svc "github.com/Jason2924/st-enginerring_test/services"

	"github.com/jszwec/csvutil"
)

func ImportProductData(pdtSrvc svc.ProductService, fileName string) {
	dir := "data/" + fileName
	file, erro := os.Open(dir)
	if erro != nil {
		panic("Failed to opening file")
	}
	reader := csv.NewReader(file)
	dcdr, erro := csvutil.NewDecoder(reader)
	if erro != nil {
		panic("Failed to decoding data")
	}
	defer file.Close()
	if dcdr != nil {
		var csvItems []ntt.ProductSchema
		erro := dcdr.Decode(&csvItems)
		if erro != nil {
			panic("Failed to parsing Table object from " + fileName)
		}
		items := make([]*ntt.ProductSchema, 0, len(csvItems))
		for _, csvItem := range csvItems {
			item := &ntt.ProductSchema{
				ID:            csvItem.ID,
				Name:          csvItem.Name,
				Price:         csvItem.Price,
				Currency:      csvItem.Currency,
				Image:         csvItem.Image,
				RatingAverage: csvItem.RatingAverage,
				RatingReviews: csvItem.RatingReviews,
			}
			items = append(items, item)
		}
		// ctrl.tableService.InsertFromFile(&items)
		pdtSrvc.InsertFromFile(items)
	}
}
