package CoinTiger_SDK_Golang

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func Sign(key, data string) string {
	hmac := hmac.New(sha512.New, []byte(key))
	hmac.Write([]byte(data))
	fmt.Println(hex.EncodeToString(hmac.Sum([]byte(""))))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}
