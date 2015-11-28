// Fyndiq API v2 implementation as described in http://developers.fyndiq.com/api-v2/
package gocommerce

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type FyndiqV2 struct {
	tr Transport
}

type FyndiqV2Settings struct {
	ProductFeedURL             *string `json:"product_feed_url"`
	ProductFeedNotificationURL *string `json:"product_feed_notification_url"`
	OrderNotificationURL       *string `json:"order_notification_url"`
}

type paginated struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type FyndiqV2ProductInfoItem struct {
	ProductID string `json:"product_id"`
	ForSale   string `json:"for_sale"`
}

type FyndiqV2ProductInfo struct {
	Results []FyndiqV2ProductInfoItem `json:"results"`
	paginated
}

type FyndiqV2OrderItem struct {
	Sku               string `json:"sku"`
	Quantity          int    `json:"quantity"`
	UnitPriceAmount   string `json:"unit_price_amount"`
	UnitPriceCurrency string `json:"unit_price_currency"`
	VatPercent        string `json:"vat_percent"`
}

type FyndiqV2Order struct {
	ID                  int                 `json:"id"`
	Created             string              `json:"created"`
	DeliveryPhone       string              `json:"delivery_phone"`
	DeliveryCompany     string              `json:"delivery_company"`
	DeliveryFirstname   string              `json:"delivery_firstname"`
	DeliveryLastname    string              `json:"delivery_lastname"`
	DeliveryAddress     string              `json:"delivery_address"`
	DeliveryCo          string              `json:"delivery_co"`
	DeliveryPostalcode  string              `json:"delivery_postalcode"`
	DeliveryCity        string              `json:"delivery_city"`
	DeliveryCountry     string              `json:"delivery_country"`
	DeliveryCountryCode string              `json:"delivery_country_code"`
	OrderRows           []FyndiqV2OrderItem `json:"order_rows"`
	DeliveryNote        string              `json:"delivery_note"`
}

type FyndiqV2Orders struct {
	Results []FyndiqV2Order `json:"results"`
	paginated
}

func NewFyndiqV2(tr Transport) *FyndiqV2 {
	return &FyndiqV2{tr}
}

// GetSettings returns the current API v2 settings
func (f *FyndiqV2) GetSettings() (*FyndiqV2Settings, error) {
	var settings FyndiqV2Settings
	var err error
	var url string
	if url, err = f.tr.URL("settings/", map[string]string{}); err != nil {
		return nil, err
	}
	var b []byte
	if b, err = f.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, settings); err != nil {
		return &settings, err
	}
	return &settings, nil
}

// UpdateSettings updates API v2 settings
func (f *FyndiqV2) UpdateSettings(settings *FyndiqV2Settings) error {
	var url string
	var err error
	if url, err = f.tr.URL("settings/", map[string]string{}); err != nil {
		return err
	}
	var b []byte
	if b, err = json.Marshal(settings); err != nil {
		return err
	}
	return f.tr.Patch(url, bytes.NewReader(b))
}

// GetProductInfo returns product status information
func (f *FyndiqV2) GetProductInfo(url string) (*FyndiqV2ProductInfo, error) {
	var productInfo FyndiqV2ProductInfo
	var err error
	var b []byte
	if b, err = f.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, productInfo); err != nil {
		return &productInfo, err
	}
	return &productInfo, nil
}

func (f *FyndiqV2) GetOrder(orderId int) (*FyndiqV2Order, error) {
	var order FyndiqV2Order
	var err error
	var url string
	if url, err = f.tr.URL("order/"+strconv.Itoa(orderId), map[string]string{}); err != nil {
		return nil, err
	}
	var b []byte
	if b, err = f.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, order); err != nil {
		return &order, err
	}
	return &order, nil
}

func (f *FyndiqV2) GetOrders(url string) (*FyndiqV2Orders, error) {
	var orders FyndiqV2Orders
	var err error
	var b []byte
	if b, err = f.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, orders); err != nil {
		return &orders, err
	}
	return &orders, nil
}
