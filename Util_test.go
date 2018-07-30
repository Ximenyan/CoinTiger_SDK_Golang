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
	InitSDK("100310001")
	str, err := Get24Hours("tchbtc")
	if err != nil {
		t.Error(str)
	}
}
func TestGetDepth(t *testing.T) {
	InitSDK("100310001")
	str, err := GetDepth("tchbtc", "step0")
	if err != nil {
		t.Error(str)
	}
}
func TestGetGetKLine(t *testing.T) {
	InitSDK("100310001")
	str, err := GetKLine("tchbtc", "1day", "160")
	if err != nil {
		t.Error(str)
	}
}
func TestGetKLineEasy(t *testing.T) {
	InitSDK("100310001")
	str, err := GetKLineEasy("tchbtc", "1day")
	if err != nil {
		t.Error(str)
	}
}
