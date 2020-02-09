// For networking with the proxies

import(fmt, net/http)

// Given IP and Port #, TCP connect with proxy
func connect(ip string, port int) string{
	url := fmt.Sprint(ip, ":", port)
	resp, err := http.Get(url)
	if err != nil{
		println("Error connecting with %s\n", url)
	}else{
		println("%s is active\n")
	}
}
