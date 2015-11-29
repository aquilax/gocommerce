package gocommerce

import (
	"github.com/aquilax/gocommerce/transport"
)

// API adapter for accessing PrestaShop API
type API struct {
	tr transport.Transport
}

// New creates new PrestaShop adapter
func New(tr transport.Transport) *API {
	return &API{tr}
}
