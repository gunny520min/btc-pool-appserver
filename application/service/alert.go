package service

import (
	"btc-pool-appserver/application/btcpoolclient"

	"github.com/gin-gonic/gin"
)

type alertHandler struct{}

var AlertService = &alertHandler{}

func (p *alertHandler) GetAlertSettings(c *gin.Context, params interface{}) (map[string]interface{}, error) {

	res := make(map[string]interface{})

	// get alert setting
	if d, err := btcpoolclient.AlertSettings(c, params); err != nil {
		return nil, err
	} else {
		res["settings"] = d
	}

	// get alert contacts
	if d, err := btcpoolclient.AlertContacts(c, params); err != nil {
		return nil, err
	} else {
		res["contactList"] = d
	}
	return res, nil
}
