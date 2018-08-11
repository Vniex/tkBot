package okex

import (
	"time"
	"testing"
	"log"
	Exchange "tkBot/exchange"
	Utils "tkBot/utils"
)

var httpClient=Utils.NewHttpClient(3,"socks5://127.0.0.1:1080")
var okexFuture = NewOKEx(httpClient, "", "")

func TestOKEx_GetDepthWithWs(t *testing.T) {
	okexFuture.GetDepthWithWs(Exchange.BTC_USD,Exchange.QUARTER_CONTRACT, func(depth *Exchange.Depth) {
		log.Print(depth)
	})
	time.Sleep(1 * time.Minute)
	okexFuture.ws.CloseWs()
}


func TestOKEx_GetTickerWithWs(t *testing.T) {
	okexFuture.GetTickerWithWs(Exchange.BTC_USD, Exchange.QUARTER_CONTRACT, func(ticker *Exchange.Ticker) {
		log.Print(ticker)
	})
	time.Sleep(1 * time.Minute)
	okexFuture.ws.CloseWs()
}