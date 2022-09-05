package modules

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"golang.org/x/net/proxy"
)

func Socks5(proxys string) {
	dialSocksProxy, err := proxy.SOCKS5("tcp", "proxy_ip", nil, proxy.Direct)
	if err != nil {
		fmt.Println(color.WhiteString("BAD - ") + color.RedString(proxys))
	}
	tr := &http.Transport{Dial: dialSocksProxy.Dial}
	proxyClient := &http.Client{
		Transport: tr,
	}
	resp, err := proxyClient.Get("https://httpbin.org/ip")
	if err == nil {
		fmt.Println(color.WhiteString("GOOD - ") + color.GreenString(proxys))
		defer resp.Body.Close()
	} else {
		fmt.Println(color.WhiteString("BAD - ") + color.RedString(proxys))
	}
}
