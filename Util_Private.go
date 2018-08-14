package CoinTiger_SDK_Golang

import (
	"encoding/json"
	"net/url"
	"strings"
)

/*创建订单*/
const CREATE_ORDER = Trading_Macro_v2 + "/order"

/*撤销订单*/
const CANCEL_ORDER = Trading_Macro_v2 + "/order/batch_cancel"

/*获取当前委托*/
const GET_NOW_ORDER = Trading_Macro_v2 + "/order/orders"

/*获取当前用户订单,成交中和未成交*/
const GET_NOW_USER_ORDER = Trading_Macro + "/order/new"

/*获取当前用户订单,成交和已撤销*/
const GET_USER_HISTORY = Trading_Macro + "/order/history"

/*用户撤单(单个)*/
const DELETE_ORDER = Trading_Macro + "/order"

/*获取资金状况*/
const GET_BALANCE = Trading_Macro + "/user/balance"

/* 创建订单				*/
/* coin_type:	交易对*/
/* price:		价格*/
/* volume:		数量*/
/* side:		买卖方向 买BUY 卖SELL*/
/* strType:		1 ：限价交易，2：市价交易*/
func CreateOder(coin_type, price, volume, side, strType string) (string, error) {

	strTime := getTimeStamp()
	u, err := url.Parse(CREATE_ORDER)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("time", strTime)
	if err != nil {
		return "", err
	}
	vals := url.Values{
		"symbol": {coin_type},
		"price":  {price},
		"volume": {volume},
		"side":   {side},
		"type":   {strType},
		"time":   {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpPostForm(u, vals)
	return str, err
}

type Orders struct {
	m map[string][]string
}

/*添加一个订单*/
/*coin_type: 交易对*/
/*order_num: 订单号*/
func (o *Orders) AddOrder(coin_type, order_num string) {
	if o.m == nil {
		o.m = make(map[string][]string)
		o.m[coin_type] = []string{order_num}
	} else if o.m[coin_type] == nil {
		o.m[coin_type] = []string{order_num}
	} else {
		o.m[coin_type] = append(o.m[coin_type], order_num)
	}
}

/*批量结束订单*/
/*o : 订单结构体*/
func CancelOrders(o Orders) (string, error) {
	strTime := getTimeStamp()
	u, err := url.Parse(CANCEL_ORDER)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("time", strTime)
	if err != nil {
		return "", err
	}
	bytes, _ := json.Marshal(o.m)
	vals := url.Values{
		"orderIdList": {string(bytes)},
		"time":        {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpPostForm(u, vals)
	return str, err
}

/*查询制定交易对的所有订单*/
/*coin_type 	: 		交易对*/
/*size		 	:		查询数量*/
func GetAllOrder(coin_type, size string) (string, error) {
	u, err := url.Parse(GET_TRADE)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("size", size)
	q.Set("time", strTime)
	q.Set("states", "new,canceled,expired,filled,part_filled,pending_cancel,")
	if err != nil {
		return "", err
	}
	vals := url.Values{
		"states": {"new,canceled,expired,filled,part_filled,pending_cancel,"},
		"symbol": {coin_type},
		"size":   {size},
		"time":   {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*查询制定交易筛选订单*/
/*coin_type 	: 		交易对*/
/*size		 	:		查询数量*/
/*states		:		查询状态*/
func GetOrders(coin_type, size, from, direct string, states, types []string) (string, error) {
	u, err := url.Parse(GET_TRADE)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("size", size)
	q.Set("from", from)
	q.Set("direct", direct)
	q.Set("time", strTime)
	strStates := ""
	strTypes := ""
	for _, v := range states {
		strStates = strStates + v + ","
	}
	for _, v := range types {
		strTypes = strTypes + v + ","
	}
	q.Set("states", strStates)
	q.Set("types", strTypes)
	if err != nil {
		return "", err
	}
	vals := url.Values{
		"from":   {from},
		"direct": {from},
		"states": {strStates},
		"types":  {strTypes},
		"symbol": {coin_type},
		"size":   {size},
		"time":   {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*获取当前用户的委托-未成交&成交中*/
func GetApprove(coin_type, offset, limit string) (string, error) {
	u, err := url.Parse(GET_NOW_USER_ORDER)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("symbol", coin_type)
	q.Set("offset", offset)
	q.Set("limit", limit)
	q.Set("time", strTime)

	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)

	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*获取用户的委托-成交&已撤销*/
func GetApproveHistory(coin_type, offset, limit string) (string, error) {
	u, err := url.Parse(GET_USER_HISTORY)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("symbol", coin_type)
	q.Set("offset", offset)
	q.Set("limit", limit)
	q.Set("time", strTime)

	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)

	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*用户撤单-单个订单*/
func DeleteOrder(coin_type, order_id string) (string, error) {
	u, err := url.Parse(DELETE_ORDER)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("symbol", coin_type)
	q.Set("order_id", order_id)
	q.Set("time", strTime)

	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)

	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpDelete(u)
	return str, err
}

/*获取资金状况(所有币种)*/
func GetAllBalance() (string, error) {
	u, err := url.Parse(GET_BALANCE)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("time", strTime)
	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*获取资金状况(单个币种)*/
func GetCoinBalance(coin_type string) (string, error) {
	u, err := url.Parse(GET_BALANCE)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("coin", coin_type)
	q.Set("time", strTime)
	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}
