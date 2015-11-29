// Fyndiq API v2 implementation as described in http://developers.fyndiq.com/api-v2/
package fyndiqv2

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type DeliveryService string

const (
	PostNord     DeliveryService = "PostNord"
	Schenker     DeliveryService = "Schenker"
	DHL          DeliveryService = "DHL"
	Bring        DeliveryService = "Bring"
	DeutschePost DeliveryService = "Deutsche Post"
	DPD          DeliveryService = "DPD"
	GLS          DeliveryService = "GLS"
	UPS          DeliveryService = "UPS"
	Hermes       DeliveryService = "Hermes"
)

type API struct {
	tr FyndiqV2Transport
}

type Settings struct {
	ProductFeedURL             *string `json:"product_feed_url"`
	ProductFeedNotificationURL *string `json:"product_feed_notification_url"`
	OrderNotificationURL       *string `json:"order_notification_url"`
}

type paginated struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type ProductInfoItem struct {
	ProductID string `json:"product_id"`
	ForSale   string `json:"for_sale"`
}

type ProductInfo struct {
	Results []ProductInfoItem `json:"results"`
	paginated
}

type OrderItem struct {
	Sku               string `json:"sku"`
	Quantity          int    `json:"quantity"`
	UnitPriceAmount   string `json:"unit_price_amount"`
	UnitPriceCurrency string `json:"unit_price_currency"`
	VatPercent        string `json:"vat_percent"`
}

type Order struct {
	ID                  int         `json:"id"`
	Created             string      `json:"created"`
	DeliveryPhone       string      `json:"delivery_phone"`
	DeliveryCompany     string      `json:"delivery_company"`
	DeliveryFirstname   string      `json:"delivery_firstname"`
	DeliveryLastname    string      `json:"delivery_lastname"`
	DeliveryAddress     string      `json:"delivery_address"`
	DeliveryCo          string      `json:"delivery_co"`
	DeliveryPostalcode  string      `json:"delivery_postalcode"`
	DeliveryCity        string      `json:"delivery_city"`
	DeliveryCountry     string      `json:"delivery_country"`
	DeliveryCountryCode string      `json:"delivery_country_code"`
	OrderRows           []OrderItem `json:"order_rows"`
	DeliveryNote        string      `json:"delivery_note"`
}

type Orders struct {
	Results []Order `json:"results"`
	paginated
}

type deliveryNoteRequest struct {
	Orders []struct {
		Order int `json:"order"`
	} `json:"orders"`
}

func New(tr FyndiqV2Transport) *API {
	return &API{tr}
}

// GetSettings returns the current API v2 settings
func (a *API) GetSettings() (*Settings, error) {
	var settings Settings
	var err error
	var url string
	if url, err = a.tr.URL("settings/", map[string]string{}); err != nil {
		return nil, err
	}
	var b []byte
	if b, err = a.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, settings); err != nil {
		return &settings, err
	}
	return &settings, nil
}

// UpdateSettings updates API v2 settings
func (a *API) UpdateSettings(settings *Settings) error {
	var url string
	var err error
	if url, err = a.tr.URL("settings/", map[string]string{}); err != nil {
		return err
	}
	var b []byte
	if b, err = json.Marshal(settings); err != nil {
		return err
	}
	return a.tr.Patch(url, bytes.NewReader(b))
}

// GetProductInfo returns product status information
func (a *API) GetProductInfo(url string) (*ProductInfo, error) {
	var productInfo ProductInfo
	var err error
	var b []byte
	if b, err = a.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, productInfo); err != nil {
		return &productInfo, err
	}
	return &productInfo, nil
}

func (a *API) GetOrder(orderId int) (*Order, error) {
	var order Order
	var err error
	var url string
	if url, err = a.tr.URL("order/"+strconv.Itoa(orderId), map[string]string{}); err != nil {
		return nil, err
	}
	var b []byte
	if b, err = a.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, order); err != nil {
		return &order, err
	}
	return &order, nil
}

func (a *API) GetOrders(url string) (*Orders, error) {
	var orders Orders
	var err error
	var b []byte
	if b, err = a.tr.Get(url); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, orders); err != nil {
		return &orders, err
	}
	return &orders, nil
}

func (a *API) GetDeliveryNotes(orderIds []int) (io.ReadCloser, error) {
	var err error
	var url string
	if url, err = a.tr.URL("delivery_notes/", map[string]string{}); err != nil {
		return nil, err
	}
	var dnr deliveryNoteRequest
	for _, orderId := range orderIds {
		dnr.Orders = append(dnr.Orders, struct {
			Order int `json:"order"`
		}{orderId})
	}
	var b []byte
	if b, err = json.Marshal(dnr); err != nil {
		return nil, err
	}
	var req *http.Request
	if req, err = a.tr.NewRequest("POST", url, bytes.NewBuffer(b)); err != nil {
		return nil, err
	}
	var resp *http.Response
	if resp, err = a.tr.Client().Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Body, nil
}
