package gocommerce

// PrestaShop is API adapter for accessing PrestaShop API
type PrestaShop struct {
	pt PrestaShopTransport
}

// NewPrestaShop creates new PrestaShop adapter
func NewPrestaShop(pt PrestaShopTransport) *PrestaShop {
	return &PrestaShop{pt}
}
