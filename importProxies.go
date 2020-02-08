//Chris Sequeira
//File imports list of Proxies
//ports them to an array
package main

//import QoQuery (JQuery for GO)
import (
	"debug/dwarf"
	"github.com/PuerkitoBio/goquery"
)

//main
func main() {
	//stores address:port
	var proxies [20]string
	//create a scrapable query from this site
	goquery.NewDocument("https://free-proxy-list.net")
	//TODO pull list
	//TODO create and populate a list
	//for each entry

}
