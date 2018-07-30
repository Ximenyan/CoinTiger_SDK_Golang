// CoinTiger_SDK_Golang project CoinTiger_SDK_Golang.go
package CoinTiger_SDK_Golang

var API_KEY string
var API_SECRET string

/*初始化SDK*/ //
func InitSDK(api_key, api_secret string) {
	API_KEY = api_key
	API_SECRET = api_secret
}
