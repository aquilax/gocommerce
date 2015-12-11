package fyndiqv1

import (
	"encoding/json"
	"github.com/aquilax/gocommerce/transport"
	"strconv"
)

type API struct {
	tr transport.Transport
}

// MetaData holds the generic meta data response header
type MetaData struct {
	Limit      int    `json: "limit"`
	Next       string `json: "next"`
	Offset     int    `json: "offset"`
	Previous   string `json: "previous"`
	TotalCount int    `json: "total_count"`
}

// MetaResponse holds generic response with meta data
type MetaResponse struct {
	Meta MetaData `json: "meta"`
}

// Variation for article group
// http://fyndiq.github.io/api-v1/#product-article-group
type Variation struct {
	ID             int    `json: "id"`
	Name           string `json: "name"`
	NumInStock     int    `json: "num_in_stock"`
	Location       string `json: "location"`
	ItemNo         string `json: "item_no"`
	PlatformItemNo string `json: "platform_item_no"`
}

// ArticleGroup for product
// http://fyndiq.github.io/api-v1/#product-article
type ArticleGroup struct {
	Name       string      `json: "name"`
	Variations []Variation `json: "variations"`
}

// Product represents single product
// http://fyndiq.github.io/api-v1/#product
type Product struct {
	Title             string       `json: "title"`
	Description       string       `json: "description"`
	Oldprice          float32      `json: "oldprice"`
	Price             float32      `json: "price"`
	MomsPercent       int          `json: "moms_percent"`
	NumInStock        int          `json: "num_in_stock"`
	State             string       `json: "state"`
	IsBlockedByFyndiq bool         `json: "is_blocked_by_fyndiq"`
	ItemNo            string       `json: "item_no"`
	PlatformItemNo    string       `json: "platform_item_no"`
	Location          string       `json: "location"`
	URL               string       `json: "url"`
	VariationGroup    ArticleGroup `json: "variation_group"`
	Images            []string     `json: "images"`
}

// ProductList represents list of products with meta data
type ProductList struct {
	MetaResponse
	Objects []Product `json: "objects"`
}

func New(tr transport.Transport) *API {
	return &API{tr}
}

// GetProducts fetches list of all products
// http://developers.fyndiq.com/api-v1/#get-read-products
func (a *API) GetProducts(url string) (*ProductList, error) {
	var productList ProductList
	var err error
	var b []byte
	if b, err = a.tr.Get(url); err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, productList)
	return &productList, err
}

// GetProduct fetches single product by ID
// http://fyndiq.github.io/api-v1/#get-read-products
func (a *API) GetProduct(id int) (*Product, error) {
	var product Product
	var err error
	var url string
	if url, err = a.tr.URL("product/"+strconv.Itoa(id), map[string]string{}); err != nil {
		return nil, err
	}
	var b []byte
	if b, err = a.tr.Get(url); err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, product)
	return &product, err
}

// DeleteProduct deletes single product by ID
// http://fyndiq.github.io/api-v1/#delete-delete-products
func (a *API) DeleteProduct(id int) error {
	var err error
	var url string
	if url, err = a.tr.URL("product/", map[string]string{}); err != nil {
		return err
	}
	return a.tr.Delete(url)
}
