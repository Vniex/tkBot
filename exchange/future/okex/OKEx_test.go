package okex

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	. "tkBot/exchange"
)

var (
	okex = NewOKEx(http.DefaultClient, "", "")
)

func TestOKEx_GetFutureDepth(t *testing.T) {
	dep, err := okex.GetFutureDepth(BTC_USD, QUARTER_CONTRACT, 5)
	assert.Nil(t, err)
	t.Log(dep.UTime)
	t.Log(dep.Pair)
	t.Log(dep.ContractType)
	t.Log(dep.AskList)
	t.Log(dep.BidList)
}


func TestOKEx_GetExchangeName(t *testing.T) {
	name:=okex.GetExchangeName()
	t.Log(name)
}

func TestOKEx_GetFee(t *testing.T) {
	fee,err:=okex.GetFee()
	t.Log(err)
	t.Log(fee)

}