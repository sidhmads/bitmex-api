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

// Individual & Bucketed Trades
type Trade struct {
	Timestamp       time.Time `json:"timestamp"`
	Symbol          string    `json:"symbol"`
	Side            string    `json:"side,omitempty"`
	Size            float32   `json:"size,omitempty"`
	Price           float64   `json:"price,omitempty"`
	TickDirection   string    `json:"tickDirection,omitempty"`
	TrdMatchID      string    `json:"trdMatchID,omitempty"`
	GrossValue      float32   `json:"grossValue,omitempty"`
	HomeNotional    float64   `json:"homeNotional,omitempty"`
	ForeignNotional float64   `json:"foreignNotional,omitempty"`
	TrdType         string    `json:"trdType,omitempty"`
}

func (trade Trade) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddTime("timestamp", trade.Timestamp)
	encoder.AddString("symbol", trade.Symbol)
	encoder.AddString("side", trade.Side)
	encoder.AddString("trdType", trade.TrdType)
	encoder.AddFloat32("size", trade.Size)
	encoder.AddFloat64("price", trade.Price)
	encoder.AddString("tickDirection", trade.TickDirection)
	encoder.AddString("trdMatchID", trade.TrdMatchID)
	encoder.AddFloat32("grossValue", trade.GrossValue)
	encoder.AddFloat64("homeNotional", trade.HomeNotional)
	encoder.AddFloat64("foreignNotional", trade.ForeignNotional)
	return nil
}
