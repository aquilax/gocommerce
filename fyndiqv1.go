package gocommerce

type FyndiqV1 struct {
	tr Transport
}

func NewFyndiqV1(tr Transport) *FyndiqV1 {
	return &FyndiqV1{tr}
}
