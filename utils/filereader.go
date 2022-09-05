package utils

import (
	"bufio"
	"os"
	"sync"

	"github.com/VertixSexyDev/goproxychecker/modules"
)

var throttle = make(chan int, 100)

func FileReader(file string, module string) {
	proxyfile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer proxyfile.Close()
	scanner := bufio.NewScanner(proxyfile)
	var wg sync.WaitGroup
	for scanner.Scan() {
		throttle <- 1
		wg.Add(1)
		if module == "1" {
			go Check(scanner.Text(), "1", &wg, throttle)
		} else if module == "2" {
			go Check(scanner.Text(), "2", &wg, throttle)
		} else if module == "3" {
			go Check(scanner.Text(), "3", &wg, throttle)
		}
	}
	wg.Wait()
}

func Check(proxy string, module string, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()
	if module == "1" {
		modules.Https(proxy)
	} else if module == "2" {
		modules.Socks4(proxy)
	} else if module == "3" {
		modules.Socks5(proxy)
	}
	<-throttle
}
