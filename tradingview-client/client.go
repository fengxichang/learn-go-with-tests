package main

import (
	"bytes"
	"fmt"
	socket "github.com/marcos-gonalons/tradingview-scraper/v2"
	"net"
	"net/http"
	neturl "net/url"
	"strings"
	"time"
)

func main() {
	tradingviewsocket, err := socket.Connect(
		func(symbol string, data *socket.QuoteData) {
			fmt.Printf("%#v", symbol)
			fmt.Printf("%#v", data)
		},
		func(err error, context string) {
			fmt.Printf("%#v", "error -> "+err.Error())
			fmt.Printf("%#v", "context -> "+context)
		},
	)
	if err != nil {
		panic("Error while initializing the trading view socket -> " + err.Error())
	}

	tradingviewsocket.AddSymbol("OANDA:EURUSD")
}

func SendHttpRequest(url, method, body string, cookies []http.Cookie, headers []map[string]string) (string, error) {
	if method == "" {
		method = "GET"
	}

	u, _ := neturl.Parse("http://127.0.0.1:8001")

	var client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(u),
			DialContext: (&net.Dialer{
				Timeout: 50 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2: true,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       50 * time.Second,
	}

	var req *http.Request

	req, _ = http.NewRequest(method, url, strings.NewReader(body))

	if nil != cookies && len(cookies) != 0 {
		for i := range cookies {
			req.AddCookie(&cookies[i])
		}
	}

	if nil != headers && len(headers) > 0 {
		for i := range headers {
			for k, v := range headers[i] {
				req.Header.Add(k, v)
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}
