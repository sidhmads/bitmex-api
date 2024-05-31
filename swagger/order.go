/*
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  [View Changelog](/app/apiChangelog)    #### Getting Started   ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](https://www.bitmex.com/app/restAPI).  *All* table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  *This is only a small subset of what is available, to get you started.*  Fill in the parameters and click the `Try it out!` button to try any of these queries.  * [Pricing Data](#!/Quote/Quote_get)  * [Trade Data](#!/Trade/Trade_get)  * [OrderBook Data](#!/OrderBook/OrderBook_getL2)  * [Settlement Data](#!/Settlement/Settlement_get)  * [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  ##### Swagger Specification  [⇩ Download Swagger JSON](swagger.json)    ## All API Endpoints  Click to expand a section.
 *
 * OpenAPI spec version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"time"

	"go.uber.org/zap/zapcore"
)

// Placement, Cancellation, Amending, and History
type Order struct {
	OrderID               string    `json:"orderID"`
	ClOrdID               string    `json:"clOrdID,omitempty"`
	ClOrdLinkID           string    `json:"clOrdLinkID,omitempty"`
	Account               float32   `json:"account,omitempty"`
	Symbol                string    `json:"symbol,omitempty"`
	Side                  string    `json:"side,omitempty"`
	SimpleOrderQty        float64   `json:"simpleOrderQty,omitempty"`
	OrderQty              float32   `json:"orderQty,omitempty"`
	Price                 float64   `json:"price,omitempty"`
	DisplayQty            float32   `json:"displayQty,omitempty"`
	StopPx                float64   `json:"stopPx,omitempty"`
	PegOffsetValue        float64   `json:"pegOffsetValue,omitempty"`
	PegPriceType          string    `json:"pegPriceType,omitempty"`
	Currency              string    `json:"currency,omitempty"`
	SettlCurrency         string    `json:"settlCurrency,omitempty"`
	OrdType               string    `json:"ordType,omitempty"`
	TimeInForce           string    `json:"timeInForce,omitempty"`
	ExecInst              string    `json:"execInst,omitempty"`
	ContingencyType       string    `json:"contingencyType,omitempty"`
	ExDestination         string    `json:"exDestination,omitempty"`
	OrdStatus             string    `json:"ordStatus,omitempty"`
	Triggered             string    `json:"triggered,omitempty"`
	WorkingIndicator      bool      `json:"workingIndicator,omitempty"`
	OrdRejReason          string    `json:"ordRejReason,omitempty"`
	SimpleLeavesQty       float64   `json:"simpleLeavesQty,omitempty"`
	LeavesQty             float32   `json:"leavesQty,omitempty"`
	SimpleCumQty          float64   `json:"simpleCumQty,omitempty"`
	CumQty                float32   `json:"cumQty,omitempty"`
	AvgPx                 float64   `json:"avgPx,omitempty"`
	MultiLegReportingType string    `json:"multiLegReportingType,omitempty"`
	Text                  string    `json:"text,omitempty"`
	TransactTime          time.Time `json:"transactTime,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

func (order Order) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("orderID", order.OrderID)
	encoder.AddString("clOrdID", order.ClOrdID)
	encoder.AddString("clOrdLinkID", order.ClOrdLinkID)
	encoder.AddString("symbol", order.Symbol)
	encoder.AddString("side", order.Side)
	encoder.AddString("pegPriceType", order.PegPriceType)
	encoder.AddString("currency", order.Currency)
	encoder.AddString("settlCurrency", order.SettlCurrency)
	encoder.AddString("ordType", order.OrdType)
	encoder.AddString("timeInForce", order.TimeInForce)
	encoder.AddString("execInst", order.ExecInst)
	encoder.AddString("contingencyType", order.ContingencyType)
	encoder.AddString("exDestination", order.ExDestination)
	encoder.AddString("ordStatus", order.OrdStatus)
	encoder.AddString("triggered", order.Triggered)
	encoder.AddString("ordRejReason", order.OrdRejReason)
	encoder.AddString("multiLegReportingType", order.MultiLegReportingType)
	encoder.AddString("text", order.Text)
	encoder.AddFloat32("account", order.Account)
	encoder.AddFloat32("orderQty", order.OrderQty)
	encoder.AddFloat32("displayQty", order.DisplayQty)
	encoder.AddFloat32("leavesQty", order.LeavesQty)
	encoder.AddFloat32("cumQty", order.CumQty)
	encoder.AddFloat64("simpleOrderQty", order.SimpleOrderQty)
	encoder.AddFloat64("price", order.Price)
	encoder.AddFloat64("stopPx", order.StopPx)
	encoder.AddFloat64("pegOffsetValue", order.PegOffsetValue)
	encoder.AddFloat64("simpleLeavesQty", order.SimpleLeavesQty)
	encoder.AddFloat64("simpleCumQty", order.SimpleCumQty)
	encoder.AddFloat64("avgPx", order.AvgPx)
	encoder.AddTime("transactTime", order.TransactTime)
	encoder.AddTime("timestamp", order.Timestamp)
	encoder.AddBool("workingIndicator", order.WorkingIndicator)
	return nil
}
