package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"os"
	"bufio"
	"time"
	"github.com/akamensky/argparse"
	"strings"
	"encoding/json"
	)

var red = "\033[1;31m"
var green = "\033[1;32m"
var yellow = "\033[1;33m"
var defcol = "\033[0m"

var re = regexp.MustCompile(`[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}(\:[\d]{2,5})?`)


// TODO: option for stdout
// var stdout = true
// if outfile is specified, this will become false

/**
* Location Struct to parse the proxy location from json 
*/
type Location struct {
	// IP string `json:"ip"`
	// City string `json:"city"`
	// Region string `json:"region"`
	// Region_code string `json:"region_code"`
	// Country string `json:"country"`
	// Country_code string `json:"country_code"`
	// Country_code_iso3 string `json:"country_code_iso3"`
	// Country_capital string `json:"country_capital"`
	// Country_tld string `json:"country_tld"`
	Country_name string `json:"country_name"`
	// Continent_code string `json:"continent_code"`
	// IN_eu bool `json:"in_eu"`
	// Postal string `json:"postal"`
	// Latitude int `json:"latitude"`
	// Longitude int `json:"longitude"`
	// Timezone string `json:"timezone"`
	// UTC_offset string `json:"utc_offset"`
	// Country_calling_code string `json:"country_calling_code"`
	// Currency string `json:"currency"`
	// Currency_name string `json:"currency_name"`
	// Languages string `json:"languages"`
	// Country_area int `json:"country_area"`
	// Country_population int `json:"country_population"`
	// ASN string `json:"asn"`
	// ORG string `json:"org"`
}

/**
 * Ascii art banner
**/
func showBanner() {
	banner := `
	_______                                                   ______                                                             
	|       \                                                 /      \                                                            
	| $$$$$$$\  ______    ______   __    __  __    __        |  $$$$$$\  _______   ______   ______    ______    ______    ______  
	| $$__/ $$ /      \  /      \ |  \  /  \|  \  |  \ ______| $$___\$$ /       \ /      \ |      \  /      \  /      \  /      \ 
	| $$    $$|  $$$$$$\|  $$$$$$\ \$$\/  $$| $$  | $$|      \\$$    \ |  $$$$$$$|  $$$$$$\ \$$$$$$\|  $$$$$$\|  $$$$$$\|  $$$$$$\
	| $$$$$$$ | $$   \$$| $$  | $$  >$$  $$ | $$  | $$ \$$$$$$_\$$$$$$\| $$      | $$   \$$/      $$| $$  | $$| $$    $$| $$   \$$
	| $$      | $$      | $$__/ $$ /  $$$$\ | $$__/ $$       |  \__| $$| $$_____ | $$     |  $$$$$$$| $$__/ $$| $$$$$$$$| $$      
	| $$      | $$       \$$    $$|  $$ \$$\ \$$    $$        \$$    $$ \$$     \| $$      \$$    $$| $$    $$ \$$     \| $$      
	 \$$       \$$        \$$$$$$  \$$   \$$ _\$$$$$$$         \$$$$$$   \$$$$$$$ \$$       \$$$$$$$| $$$$$$$   \$$$$$$$ \$$      
											|  \__| $$                                              | $$                          
											 \$$    $$                                              | $$                          
											  \$$$$$$                                                \$$                          
	`
	fmt.Println(banner)
}

/**
 * Parameter: string url
 * Description: Checks if a string is a url before making a request
**/
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

/**
* Parameter: IP
* Description: Makes a request to "https://ipapi.co" and parse the json result to get the
* 			   location of the proxy
**/
func getlocation(ip string) string {
	tempurl, err := http.Get("https://ipapi.co/" + ip + "/json/")
	if err != nil {
		fmt.Println(err)
	}
	defer tempurl.Body.Close()
	body, _ := ioutil.ReadAll(tempurl.Body)
	// fmt.Println(string(body))

	var location Location
	json.Unmarshal(body, &location)
	return location.Country_name
}

func warning(msg string) {

}

func info(msg string) string{
	return red + "[" + green + "+" + red + "] - " + defcol + msg
}

func error(msg string) string {
	return red + "[" + yellow + "!" + red + "] - " + defcol + msg
}

/** 
* Parameter: url --> string, outfile --> pointer to a file
* Description: Makes a proxy website and use regex to get all the IP and its port(if available)
*			   from the site and saves them to a file
**/
func scraper(url string, outfile *os.File) {
	// Split result to good proxy and bad proxy
	resp, err := http.Get(url) 

	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Status Code: ", resp.StatusCode)
	// Checks the status of the page
	if resp.StatusCode == 404 {
		fmt.Println(error("Page doesn't exist"))
	}else{
		results := re.FindAllString(string(body), -1)
		if len(results) == 0 {
			fmt.Println(error("No proxy was found on "), url)
		}else {
			fmt.Printf(info("Found %d on %s\n"), len(results), url)
			// write to a file
			for _, proxy := range results {
				splitProxy := strings.Split(proxy, ":")
				location := getlocation(splitProxy[0])
				fmt.Printf("Current proxy: %s, location: %s\n", proxy, location)
				_, err := outfile.WriteString(proxy + "--> " + location + "\n")
				if err != nil {
					fmt.Println("An error occured ", err)
					outfile.Close()
				}
			}
		}
	}
}

/**
* Parameter: url --> string
* Description: Test if the proxy works by making a request to google.com. if the request code is 200
* 			   then we know it is valid. If otherwise, then we discard it.
**/
func isAlive(url string) bool {
	// Need to set HTTP_PROXY 

	var PTransport = &http.Transport { Proxy: http.ProxyFromEnvironment }
	client	:= http.Client { Transport: PTransport }

	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	}

	return false
}

/**
* Description: Entry point of the program
*/
func main(){
	// 1) read file from users -- DONE
	// 2) For each line in the file(verify if it is a url) -- DONE
	// 2A) CHECK IF URL IS "ALIVE" 
	// 3) Create a thread for each line -- DONE
	// 4) then run the scraper -- DONE


	// Argument Parsing
	parser := argparse.NewParser("Proxy Scraper", "Proxy Scraper implemented in golang. By PabloPotat0")
	inputfile := parser.String("i", "inputfile", &argparse.Options{Help: "Proxy websites to scrap"})
	outfile := parser.String("o", "outfile", &argparse.Options{Help: "Good proxies will be stored here"})
	check := parser.Flag("", "check", &argparse.Options{Help: "status of the proxy, alive or dead"})
	chain := parser.String("", "chain", &argparse.Options{Help: "linked list of proxy to direct traffic"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	
	// error checking
	if *inputfile == "" {
		fmt.Println(error("input file not specified. Using default file: proxylist.txt"))
		*inputfile = "proxylist.txt"
	} 
	
	if *outfile == ""{
		fmt.Println(error("output file not specified. Using default file: good_proxy.txt"))
		*outfile = "good_proxy.txt"
	}

	fmt.Println(*check)
	if *check {
		fmt.Println("what am in?? ", *check)
		//call the isAlive()
	}

	if *chain == "" {
		fmt.Println("Under construction ")
	}

	// Create outfile
	good, _ := os.Create(*outfile)

	// Banner - duh
	showBanner()

	fmt.Println("Let's begin scraping....")

	file, err := os.Open(*inputfile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Url validator
		if isValidUrl(scanner.Text()) {
			go scraper(scanner.Text(), good)
			time.Sleep(time.Second)
		}else{
			fmt.Printf(error("%s is not a valid url\n"), scanner.Text())
		}
	}

	good.Close()
	fmt.Println("Ending")
}

