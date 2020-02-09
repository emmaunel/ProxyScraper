// For networking with the proxies

import(fmt, net/http)

// Given IP and Port #, TCP connect with proxy
// Return false if connection fails, else return true
func connect(ip string, port int) string{
	url := fmt.Sprint(ip, ":", port)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		println("Error connecting with %s\n", url)
		return false
	}else{
		println("%s is active\n")
		return true
	}
}
