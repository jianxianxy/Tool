package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
http://www.zhilianhttp.com/getapi.html
https://www.kuaidaili.com/free/inha/
60.205.229.126
211.159.171.58
*/
func main() {
	//uri := "http://www.net.cn/static/customercare/yourip.asp"
	uri := "http://127.0.0.1/test/ip.php"

	agent := "local"
	//agent := "http://183.166.110.58:33962"
	rep := httpReq(uri, agent, "post")
	fmt.Println(string(rep))
	os.Exit(1)
}

func httpReq(uri string, agent string, method string) (body []byte) {
	request, _ := http.NewRequest("GET", uri, nil)
	if method == "post" {
		fdata := url.Values{}
		fdata.Set("phone", "18212345678")
		fdata.Set("type", "phone_login")
		request, _ = http.NewRequest("POST", uri, strings.NewReader(fdata.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Set("Cookie", "name=anny")
		request.Header.Set("Content-Length", strconv.Itoa(len(fdata.Encode())))
	}
	//随机返回User-Agent 信息
	request.Header.Set("User-Agent", getAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	proxy, err := url.Parse(agent)
	if err != nil {
		fmt.Printf("代理错误:%s\n", err)
		os.Exit(1)
	}
	//设置超时时间
	timeout := time.Duration(30 * time.Second)
	client := &http.Client{}
	if agent != "local" {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: timeout,
		}
		fmt.Printf("使用代理:%s\n", proxy)
	}
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		fmt.Printf("Error:func @ httpReq %s\n", err)
		os.Exit(1)
	} else {
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				body = append(body, buf...)
			}
		default:
			body, _ = ioutil.ReadAll(response.Body)

		}
	}
	return
}

/**
* 随机返回一个User-Agent
 */
func getAgent() string {
	agent := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(agent)
	return agent[r.Intn(len)]
}
