package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"log"
	"bufio"
)

type Proxy struct {
	IP string

}

func main(){
	var proxyscape = "https://api.proxyscrape.com/?request=getproxies&proxytype=http&timeout=5000&country=US&anonymity=elite&ssl=yes"
	// var test = "https://api.getproxylist.com/proxy"
	fmt.Println("Let's begin scraping")
	res, err := http.Get(proxyscape)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
        log.Fatal(err)
	}
	
	fmt.Println(string(resBody))
	filename, err := os.Create("proxylist.txt")
	defer filename.Close()

	file := bufio.NewWriter(filename)
	content, _ := file.WriteString(string(resBody))
	fmt.Printf("Wrote %d bytes\n", content)

	file.Flush()
	
}