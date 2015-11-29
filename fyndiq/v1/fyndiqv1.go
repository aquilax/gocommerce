package fyndiqv1

import (
	"github.com/aquilax/gocommerce/transport"
)

type API struct {
	tr transport.Transport
}

func New(tr transport.Transport) *API {
	return &API{tr}
}
