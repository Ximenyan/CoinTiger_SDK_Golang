package CoinTiger_SDK_Golang

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const API_PATH = "https://api.cointiger.pro"
const Trading_Macro_v2 = API_PATH + "/exchange/trading/api/v2"
const Market_Macro = API_PATH + "/exchange/trading/api/market"

/*获取系统时间*/
const GET_TIMESTAMP = Trading_Macro_v2 + "/timestampw"

/*获取支持的所有币种*/
const GET_CURRENCYS = Trading_Macro_v2 + "/currencys"

/*获取支持24小时行情*/
const GET_24HOURS = Market_Macro + "/detail"

/*获取深度盘口*/
const GET_DEPTH = Market_Macro + "/depth"

/*获取历史K线*/
const GET_KLine = Market_Macro + "/history/kline"

/*获取成交历史数据*/
const GET_TRADE = Market_Macro + "/history/trade"

/*HTTP GET*/
func httpGet(url *url.URL) (string, error) {
	fmt.Println(url.String())
	resp, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	nbytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(nbytes), nil
}

/*获取服务器时间*/
func GetTimeStamp() (string, error) {
	u, err := url.Parse(GET_TIMESTAMP)
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取支持的币种*/
func GetCurrencys() (string, error) {
	u, err := url.Parse(GET_CURRENCYS)
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取24小时行情*/
func Get24Hours(coin_type string) (string, error) {
	u, err := url.Parse(GET_24HOURS)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/**
*	获取深度
*   coin_type:交易对
*	depth_type:Depth 类型 :"step0", "step1", "step2"
**/
func GetDepth(coin_type, depth_type string) (string, error) {
	u, err := url.Parse(GET_DEPTH)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("type", depth_type)
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取历史K线*/
/*period: K线类型: 1min,5min,15min,30min,60min,1day,1week,1month*/
/*size: 获取数量: [1,2000]*/
func GetKLine(coin_type, period, size string) (string, error) {
	u, err := url.Parse(GET_KLine)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("period", period)
	q.Set("size", size)
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取历史K线*/
/*period: K线类型: 1min,5min,15min,30min,60min,1day,1week,1month*/
func GetKLineEasy(coin_type, period string) (string, error) {
	u, err := url.Parse(GET_KLine)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("period", period)
	q.Set("size", "200")
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}
