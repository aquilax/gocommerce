package gocommerce

type FyndiqV2 struct {
	tr Transport
}

func NewFyndiqV2(tr Transport) *FyndiqV2 {
	return &FyndiqV2{tr}
}
