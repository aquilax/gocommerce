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

type OrderPackage struct {
	Service  DeliveryService `json:"service "`
	Tracking string          `json:"tracking"`
	Sku      []string        `json:"sku"`
}

type OrderPackages struct {
	Packages []OrderPackage `json:"packages"`
}

type BulkOrderPackage struct {
	Order string `json:"order"`
	OrderPackage
}

type BulkOrderPackages struct {
	Packages []BulkOrderPackage `json:"packages"`
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
	err = json.Unmarshal(b, settings)
	return &settings, err
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
	err = json.Unmarshal(b, productInfo)
	return &productInfo, err
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
	err = json.Unmarshal(b, order)
	return &order, err
}

func (a *API) GetOrders(url string) (*Orders, error) {
	var orders Orders
	var err error
	var b []byte
	if b, err = a.tr.Get(url); err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, orders)
	return &orders, err
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
	resp, err = a.tr.Client().Do(req)
	defer resp.Body.Close()
	return resp.Body, err
}

func (a *API) SetOrderPackages(orderId int, packages *OrderPackages) error {
	var url string
	var err error
	if url, err = a.tr.URL("packages/"+strconv.Itoa(orderId), map[string]string{}); err != nil {
		return err
	}
	var b []byte
	if b, err = json.Marshal(packages); err != nil {
		return err
	}
	return a.tr.Post(url, bytes.NewReader(b))
}

func (a *API) SetBulkPackages(packages *BulkOrderPackages) error {
	var url string
	var err error
	if url, err = a.tr.URL("packages/", map[string]string{}); err != nil {
		return err
	}
	var b []byte
	if b, err = json.Marshal(packages); err != nil {
		return err
	}
	return a.tr.Post(url, bytes.NewReader(b))
}
