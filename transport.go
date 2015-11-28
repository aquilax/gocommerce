package gocommerce

// Transport is general transport interface
type Transport interface {
	Get(url string) ([]byte, error)
}
