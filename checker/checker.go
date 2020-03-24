package checker

import(
	"time"
	"net/http"
	"net/url"
)

func Checker(proxy string, port string) (res *http.Response, err error) {
	timeout := time.Duration(2 * time.Second)
	proxyUrl, err := url.Parse("http://" + proxy + ":" + port)
	reqUrl, err := url.Parse("https://ispycode.com/web/hello.html")

	transpot := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{
		Timeout: timeout,
		Transport: transpot,
	}

	res, err = client.Get(reqUrl.String())

	return res, err
}

// This function is really slow. maybe I have to set a timeout 
// TODO: Come back to this
// func isAlive(protocol string, ip string, port string){
// 	proxyStr := protocol + "://" + ip + ":" + port
// 	proxyUrl, err := url.Parse(proxyStr)
// 	fmt.Println(proxyUrl)
// 	if err != nil {
// 		panic(err)
// 	}

// 	urlStr := "https://ispycode.com/web/hello.html"
// 	// urlStr := "https://api.ipify.org/"

// 	//adding the proxy settings to the Transport object
// 	transpot := &http.Transport { Proxy: http.ProxyURL(proxyUrl),
// 								// TLSClientConfig: &tls.Config{},
// 	 }
// 	//adding the Transport object to the http Client
// 	client	:= &http.Client { Transport: transpot,
// 							// Timeout:   time.Duration(2 * time.Second),
// 						}

// 	//generating the HTTP GET request
// 	request, err := http.NewRequest("GET", urlStr, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//printing the request to the console
// 	// dump, _ := httputil.DumpRequest(request, false)
// 	// fmt.Println(string(dump))

// 	//calling the URL
// 	response, err := client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}


// 	fmt.Println(response.StatusCode)
// 	// data, _ := ioutil.ReadAll(response.Body)
// 	// fmt.Println("Alive: ", string(data))
// }