// Fyndiq API v2 implementation as described in http://developers.fyndiq.com/api-v2/
package gocommerce

import (
	"bytes"
	"encoding/json"
)

type FyndiqV2 struct {
	tr Transport
}

type FyndiqV2Settings struct {
	ProductFeedURL             *string `json:"product_feed_url"`
	ProductFeedNotificationURL *string `json:"product_feed_notification_url"`
	OrderNotificationURL       *string `json:"order_notification_url"`
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
