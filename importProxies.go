//Chris Sequeira
//File imports list of Proxies
//ports them to an array
package main

//import QoQuery (JQuery for GO)
import (
	"bufio"
	"debug/dwarf"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
)

//main
func main() {
	//stores address:port
	var proxies []String

	//TODO Make HTTP Get request with link
	//TODO HOW DOES THIS RETURN
	content, badOut := http.Get("https://api.proxyscrape.com/?request=displayproxies&proxytype=all&timeout=7000&country=all&anonymity=elite&ssl=yes")

	//error handling
	if badOut != nil {
		log.Fatal(badOut)
	}

	//TODO HOW DOES THIS WORK
	defer content.body.close()

	//create a scanner object
	scanner := bufio.NewScanner(page)
	scanner.Split(bufio.ScanLines)

	//scan lines into array
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	//test proxies

	//organize proxies by testing results

	//output

}
