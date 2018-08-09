package okex

import (
	"time"
	"net/http"
	"testing"
	"log"
	. "tkBot/exchange"
)

var okexFuture = NewOKEx(http.DefaultClient, "", "")

func TestOKEx_GetDepthWithWs(t *testing.T) {
	okexFuture.GetDepthWithWs(BTC_USD, QUARTER_CONTRACT, func(depth *Depth) {
		log.Print(depth)
	})
	time.Sleep(1 * time.Minute)
	okexFuture.ws.CloseWs()
}


func TestOKEx_GetTickerWithWs(t *testing.T) {
	okexFuture.GetTickerWithWs(BTC_USD, QUARTER_CONTRACT, func(ticker *Ticker) {
		log.Print(ticker)
	})
	time.Sleep(1 * time.Minute)
	okexFuture.ws.CloseWs()
}