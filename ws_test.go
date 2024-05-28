package bitmex

import (
	"fmt"
	"log"
	"testing"

	"github.com/sidhmads/bitmex-api"
	"github.com/sidhmads/bitmex-api/swagger"
)

func testBitMeX() *BitMEX {
	b := New(nil, HostTestnet, "", "", true)
	//b.SetHttpProxy("http://127.0.0.1:1081")
	return b
}

func TestBitMEXConnect(t *testing.T) {
	b := testBitMeX()
	subscribeInfos := []SubscribeInfo{
		{Op: BitmexWSOrderBookL2, Param: "XBTUSD"},
	}
	err := b.Subscribe(subscribeInfos)
	if err != nil {
		log.Fatal(err)
	}

	b.On(BitmexWSOrderBookL2, func(ob OrderBookDataL2, symbol string) {
		m := ob.OrderBook()
		fmt.Printf("\rOrderbook Asks: %#v Bids: %#v                            ", m.Asks[0], m.Bids[0])
	})

	b.StartWS()

	select {}
}

func TestBitMEXWS(t *testing.T) {
	b := testBitMeX()
	subscribeInfos := []SubscribeInfo{
		{Op: BitmexWSOrderBookL2, Param: "XBTUSD"},
	}
	err := b.Subscribe(subscribeInfos)
	if err != nil {
		log.Fatal(err)
	}

	b.On(BitmexWSOrderBookL2, func(ob OrderBookDataL2, symbol string) {
		m := ob.OrderBook()
		fmt.Printf("\rOrderbook Asks: %#v Bids: %#v                            ", m.Asks[0], m.Bids[0])
	})

	b.StartWS()

	select {}
}

func TestBitMEX_Subscribe(t *testing.T) {
	b := testBitMeX()
	subscribeInfos := []SubscribeInfo{
		{Op: BitmexWSOrderBookL2_25, Param: "XBTUSD"},
	}
	err := b.Subscribe(subscribeInfos)
	if err != nil {
		log.Fatal(err)
	}

	b.On(BitmexWSOrderBookL2_25, func(ob OrderBookDataL2, symbol string) {
		m := ob.OrderBook()
		fmt.Printf("\rOrderbook Asks: %#v Bids: %#v                            ", m.Asks[0], m.Bids[0])
	})

	b.StartWS()

	select {}
}

func TestBitMEX_Subscribe_Various(t *testing.T) {
	b := testBitMeX()
	subscribeInfos := []SubscribeInfo{
		{Op: BitmexWSQuote, Param: "XBTUSD"},
		{Op: BitmexWSTrade, Param: "XBTUSD"},
		{Op: BitmexWSOrderBookL2_25, Param: "XBTUSD"},
		{Op: BitmexWSQuote, Param: "XRPUSD"},
		{Op: BitmexWSTrade, Param: "XRPUSD"},
		{Op: BitmexWSOrderBookL2_25, Param: "XRPUSD"},
	}
	err := b.Subscribe(subscribeInfos)
	if err != nil {
		log.Fatal(err)
	}

	b.On(bitmex.BitmexWSQuote, func(quotes []*swagger.Quote, action string) {
		fmt.Printf("\rQuote action: %v, Ask: %#v Bid: %#v                            ", action, quotes[0].AskPrice, quotes[0].BidPrice)
	}).On(bitmex.BitmexWSTrade, func(trades []*swagger.Trade, action string) {
		fmt.Printf("\rTrade action: %v, Price: %#v Amount: %#v                            ", action, trades[0].Price, trades[0].Size)
	}).On(BitmexWSOrderBookL2_25, func(ob OrderBookDataL2, symbol string) {
		m := ob.OrderBook()
		fmt.Printf("\rOrderbook symbol: %v, Asks: %#v Bids: %#v                            ", symbol, m.Asks[0], m.Bids[0])
	})

	b.StartWS()

	select {}
}
