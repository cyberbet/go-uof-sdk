package api

import (
	"github.com/minus5/go-uof-sdk"
)

const (
	probabilitiesEvent  = "/v1/probabilities/{{.EventURN}}"
	probabilitiesMarket = "/v1/probabilities/{{.EventURN}}/{{.MarketID}}"
)

// Get probabilities for a sport event.
func (a *API) ProbabilitiesEvent(eventURN uof.URN) (*uof.Cashout, error) {
	var c uof.Cashout
	return &c, a.getAs(&c, probabilitiesEvent, &params{EventURN: eventURN})
}

// Get probabilities for a sport event's specific market.
func (a *API) ProbabilitiesEventMarket(eventURN uof.URN, marketID int) (*uof.Cashout, error) {
	var c uof.Cashout
	return &c, a.getAs(&c, probabilitiesMarket, &params{EventURN: eventURN, MarketID: marketID})
}
