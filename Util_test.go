package CoinTiger_SDK_Golang

import (
	"testing"
)

func TestGetTimeStamp(t *testing.T) {
	str, err := GetTimeStamp()
	if err != nil {
		t.Error(str)
	}
}

func TestGetCurrencys(t *testing.T) {
	str, err := GetCurrencys()
	if err != nil {
		t.Error(str)
	}
}

func TestGet24Hours(t *testing.T) {
	InitSDK("100310001", "NGMzMjI3MzlhNjYyOTE4NWY34234GDFJRTvcwertyuytQWEREWQQ3234mZDA0MzhmZWZmNg==")
	str, err := Get24Hours("tchbtc")
	if err != nil {
		t.Error(str)
	}
}
func TestGetDepth(t *testing.T) {
	InitSDK("100310001", "NGMzMjI3MzlhNjYyOTE4NWY34234GDFJRTvcwertyuytQWEREWQQ3234mZDA0MzhmZWZmNg==")
	str, err := GetDepth("tchbtc", "step0")
	if err != nil {
		t.Error(str)
	}
}
func TestGetGetKLine(t *testing.T) {
	InitSDK("100310001", "NGMzMjI3MzlhNjYyOTE4NWY34234GDFJRTvcwertyuytQWEREWQQ3234mZDA0MzhmZWZmNg==")
	str, err := GetKLine("tchbtc", "1day", "160")
	if err != nil {
		t.Error(str)
	}
}

func TestGetTrade(t *testing.T) {
	InitSDK("100310001", "NGMzMjI3MzlhNjYyOTE4NWY34234GDFJRTvcwertyuytQWEREWQQ3234mZDA0MzhmZWZmNg==")
	str, err := GetTrade("tchbtc", "10")
	if err != nil {
		t.Error(str)
	}
}
func TestSign(t *testing.T) {
	Sign("api_key100310001", "NGMzMjI3MzlhNjYyOTE4NWY34234GDFJRTvcwertyuytQWEREWQQ3234mZDA0MzhmZWZmNg==")
}
func TestCreateOder(t *testing.T) {
	InitSDK("100310001", "NGMzMjI3MzlhNjYyOTE4NWY34234GDFJRTvcwertyuytQWEREWQQ3234mZDA0MzhmZWZmNg==")
	CreateOder("tchbtc", "10.1", "100.1", "BUY", "1")
}
func TestCancelOder(t *testing.T) {
	InitSDK("100310001", "NGMzMjI3MzlhNjYyOTE4NWY34234GDFJRTvcwertyuytQWEREWQQ3234mZDA0MzhmZWZmNg==")
	o := Orders{}
	o.AddOrder("tchbtc", "1234")
	CancelOrders(o)
}
