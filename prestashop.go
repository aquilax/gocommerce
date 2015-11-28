package gocommerce

// PrestaShop is API adapter for accessing PrestaShop API
type PrestaShop struct {
	tr Transport
}

// NewPrestaShop creates new PrestaShop adapter
func NewPrestaShop(tr Transport) *PrestaShop {
	return &PrestaShop{tr}
}
