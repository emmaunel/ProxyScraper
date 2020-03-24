package core

import (
	color "../colors"
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	// "net/http/httputil"
	// "crypto/tls"
	"regexp"
	"time"
)

const timeout time.Duration = 10

var httpRe = regexp.MustCompile(`([\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}).*?(\d{1,5}).*?([A-Z][A-Z]).*?(\bno).*?(\bno|\byes)`)
var httpsRe = regexp.MustCompile(`([\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}).*?(\d{1,5}).*?([A-Z][A-Z]).*?(\bno).*?(\bno|\byes)`)
var socksRe = regexp.MustCompile(`([\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}).*?(\d{1,5}).*?(Socks.?).*?`)
var ip string
var port string
var country string

func showStatus() {

}

/// Proxy chain option.
/// Name: Proxy scraper/Chain
//// Desrciption: Find a byunch of proxy and creates a proxy chaing between them and direct your chain between them.

// https://stackoverflow.com/questions/40328025/tcp-connection-over-tor-in-golang
// Running the scraper through tor


func Http_proxies(check bool){
	resp, err := http.Get("https://free-proxy-list.net")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// Regex in go sucks. It doesn't parse regex group well
	// This is the best I could come up with
	results := httpRe.FindAllStringSubmatch(string(body), -1)

	/**
	 * Example of a return regex result
	 * [18.163.28.22</td><td>1080</td><td>HK</td><td class='hm'>Hong Kong</td><td>anonymous</td><td class='hm'>no</td><td class='hx'>no 18.163.28.22 1080 HK no no]
	 * Yea. The zero index is the whole string then next to it is the actual data.
	 * Why am I doing this???
	 * The number below correspond to the data in the array above
	 * I spend too much time on this
	 **/
	for _, proxy := range results {
			for i, j := range proxy {
				// TODO: Check if the last element in the array is yes
				// yes indicates it is a https. We don't want that in this function
				if i == 1{
					ip = j
				} else if i == 2{
					port = j
				} else if i == 3 {
					country = j
				}
				// fmt.Println(j)
			}
			fmt.Println(color.PrintProxy(ip, port, "HTTP"))
			fmt.Println("Location: " + country)
			if check {
				fmt.Println(check)
			}
			time.Sleep(time.Second / 2) // Uncomment if you want to be fast

	}
}

func HttpsProxies(check bool){
	resp, err := http.Get("https://www.sslproxies.org")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)


	// Regex in go sucks. It doesn't parse regex group well
	// This is the best I could come up with
	results := httpsRe.FindAllStringSubmatch(string(body), -1)

	for _, proxy := range results {
		for i, j := range proxy {
			if i == 1{
				ip = j
			} else if i == 2{
				port = j
			} 
	// 		// fmt.Println(j)
		}
		fmt.Println(color.PrintProxy(ip, port, "HTTPS"))
		time.Sleep(time.Second / 2) // Uncomment if you want to be fast
		if check {
			fmt.Println(check)
			// isAlive("https", ip, port)
		}
	}
}

func SocksProxies(){
	resp, err := http.Get("https://www.socks-proxy.net")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)


	// Regex in go sucks. It doesn't parse regex group well
	// This is the best I could come up with
	results := socksRe.FindAllStringSubmatch(string(body), -1)

	for _, proxy := range results {
		for i, j := range proxy {
			if i == 1{
				ip = j
			} else if i == 2{
				port = j
			} 
	// 		// fmt.Println(j)
		}
		fmt.Println(color.PrintProxy(ip, port, "Sock4"))
		time.Sleep(time.Second / 2) // Uncomment if you want to be fast
	}
}

