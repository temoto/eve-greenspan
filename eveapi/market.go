package eveapi

import (
	"encoding/xml"
)

// <row orderID="639587440" charID="118406849" stationID="60003760" volEntered="25"
//   volRemaining="4" minVolume="1" orderState="0" typeID="26082" range="32767"
//   accountKey="1000" duration="1" escrow="0.00" price="3399999.98" bid="0"
//   issued="2008-02-03 22:35:54"/>
type MarketOrder struct {
	XMLName xml.Name `xml:"row"`

	OrderID      int `xml:"orderID,attr"`      // Unique order ID for this order. Note that these are not guaranteed to be unique forever, they can recycle. But they are unique for the purpose of one data pull.
	CharID       int `xml:"charID,attr"`       // ID of the character that physically placed this order.
	StationID    int `xml:"stationID,attr"`    // ID of the station the order was placed in.
	VolEntered   int `xml:"volEntered,attr"`   // Quantity of items required/offered to begin with.
	VolRemaining int `xml:"volRemaining,attr"` // Quantity of items still for sale or still desired.
	MinVolume    int `xml:"minVolume,attr"`    // For bids (buy orders), the minimum quantity that must be sold in one sale in order to be accepted by this order.

	// As of 2011-09-08 MarketOrders.xml.aspx will now return all active orders plus all orders issued in the last 7 days. An optional "orderID" parameter can be provided to fetch any order belonging to your character/corporation.
	State byte `xml:"orderState,attr"` // Valid states: 0 = open/active, 1 = closed, 2 = expired (or fulfilled), 3 = cancelled, 4 = pending, 5 = character deleted. Note: As of Incarna 1.0, the API will no longer return orders which have an orderState other than 0. Thus, no non-open orders will be returned.

	TypeID      int64   `xml:"typeID,attr"`     // ID of the type (references the invTypes table) of the items this order is buying/selling.
	Range       int     `xml:"range,attr"`      // The range this order is good for. For sell orders, this is always 32767. For buy orders, allowed values are: -1 = station, 0 = solar system, 5/10/20/40 Jumps, 32767 = region.
	AccountKey  int     `xml:"accountKey,attr"` // Which division this order is using as its account. Always 1000 for characters, but in the range 1000 to 1006 for corporations.
	Duration    int     `xml:"duration,attr"`   // How many days this order is good for. Expiration is issued + duration in days.
	ParseEscrow string  `xml:"escrow,attr"`
	ParsePrice  string  `xml:"price,attr"`
	Escrow      ISK     // How much ISK is in escrow. Valid for buy orders only (I believe).
	Price       ISK     // The cost per unit for this order.
	Bid         bool    `xml:"bid,attr"` // If true, this order is a bid (buy order). Else, sell order.
	ParseIssued string  `xml:"issued,attr"`
	Issued      EveTime // When this order was issued.
}

type xmlRowset struct {
	XMLName xml.Name `xml:"rowset"`
}

type MarketOrdersResponse struct {
	baseResponse
}

func (c *Client) CharMarketOrders() (*MarketOrdersResponse, error) {
	return nil, nil
}
