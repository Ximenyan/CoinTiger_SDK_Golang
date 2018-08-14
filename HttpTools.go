package CoinTiger_SDK_Golang

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*获取时间戳*/
func getTimeStamp() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[:13]
}

/*HTTP GET*/
func httpGet(url *url.URL) (string, error) {
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

/* HTTP POST*/
func httpPostForm(url *url.URL, values url.Values) (string, error) {
	//http.MethodDelete
	resp, err := http.PostForm(url.String(), values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/* HTTP Delete*/
func httpDelete(url *url.URL) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, url.String(), nil)
	if err != nil {
		return "", err
	}

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
