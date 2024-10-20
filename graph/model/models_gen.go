// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewProduct struct {
	Name   string     `json:"name"`
	Price  string     `json:"price"`
	Image  string     `json:"image"`
	Rating *NewRating `json:"rating"`
}

type NewRating struct {
	Average float64 `json:"average"`
	Reviews int     `json:"reviews"`
}

type Product struct {
	ID     int  `json:"id"`
	Name   string  `json:"name"`
	Price  string  `json:"price"`
	Image  string  `json:"image"`
	Rating *Rating `json:"rating"`
}

type Query struct {
}

type Rating struct {
	Average float64 `json:"average"`
	Reviews int     `json:"reviews"`
}