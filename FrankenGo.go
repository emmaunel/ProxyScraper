package main

import (
	"fmt"
	"go/scanner"
	"go/types"
	"net/http"
	"os"
	"io/ioutil"
	"log"
	"bufio"
	"strings"
)

type Proxy struct {
	IP string

}

func getProxies() Slice {
	//HTTP GET list of proxies in IP:PORT format
	var proxyscape = "https://api.proxyscrape.com/?request=getproxies&proxytype=http&timeout=5000&country=US&anonymity=elite&ssl=yes"
	// var test = "https://api.getproxylist.com/proxy"
	fmt.Println("Let's begin scraping")
	res, err := http.Get(proxyscape)

	//error checking
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//reads all returned lists
	resBody, err := ioutil.ReadAll(res.Body)

	//error checking for local read
	if err != nil {
		log.Fatal(err)
	}
	return resBody
}

func writeProxies(Slice resBody) types.Nil {
	//print results to screen
	fmt.Println(string(resBody))
	//write overwritingly to text file
	filename, err := os.Create("proxylist.txt")
	defer filename.Close()

	//writing to a file
	file := bufio.NewWriter(filename)
	content, _ := file.WriteString(string(resBody))
	fmt.Printf("Wrote %d bytes\n", content)

	//error checking
	if err != nil {
		log.Fatal(err)
	}

	//clear buffer
	file.Flush()

	return types.Nil{}
}

func mapProxies() map[]string {
	//readline
	//split by colon
	//pass tuple
	listFile, errorMsg := os.Open("proxylist.txt")
	if errorMsg != nil {
		log.Fatal(errorMsg)
	}
	defer listFile.Close()

	//create and initialize an IP:port dictionary
	proxies =  make(map[string]int)

	//populates map with values
	scan := bufio.NewScanner(listFile)
	for scan.Scan() {
		var sliced = strings.split(scanner.Text(), ":")
		proxies[sliced[0]]=sliced[1]
	}

	//additional error checking
	if errorMsg := scanner.Err(); err != nil {
		log.Fatal(errorMsg)
	}
	return proxies
}

func testProxies(slice proxies) Slice[]Boolean{
	//returns boolean results for each entry in the proxy list
	//create slice of booleans
	//for each entry in proxies:
		//test
		//if active:
			//append True to  Slice
		//else:
			//append False to Slice
	return //slice
}

func outputProxies(Slice proxies, Slice proxyResults){
	//for each entry in proxyResults:
		//if entry == False:
			//proxies[i] = Type.Null

}

func connect()#Haacked by potato

func main(){
	writeProxies(getProxies())
	mapProxies()
	//outputProxies(testProxies())



}
