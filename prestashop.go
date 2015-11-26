package gocommerce

type PrestaShop struct {
	pt PrestaShopTransport
}

func NewPrestaShop(pt PrestaShopTransport) *PrestaShop {
	return &PrestaShop{pt}
}
