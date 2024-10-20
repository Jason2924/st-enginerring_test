package graph

import (
	svc "github.com/Jason2924/st-enginerring_test/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductService svc.ProductService
}
