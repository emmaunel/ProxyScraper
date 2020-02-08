//Chris Sequeira
//File imports list of Proxies
//ports them to an array
package main

//import QoQuery (JQuery for GO)
import (
	"github.com/PuerkitoBio/goquery"
)

//main
func main() {
	//create a scrapable query from this site
	goquery.NewDocument("https://free-proxy-list.net")
	//TODO pull list
	//TODO createand populate a list

}
