package modules

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)

func Https(proxy string) {
	proxyUrl, parseerr := url.Parse("http://" + proxy)
	if parseerr == nil {
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		proxyClient := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
		resp, err := proxyClient.Get("https://httpbin.org/ip")
		if err == nil {
			fmt.Println(color.WhiteString("GOOD - ") + color.GreenString(proxy))
			defer resp.Body.Close()
		} else {
			fmt.Println(color.WhiteString("BAD - ") + color.RedString(proxy))
		}
	} else {
		fmt.Println(color.WhiteString("BAD - ") + color.RedString(proxy))
	}
}
