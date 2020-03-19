package main

import (
	test "./core"
	color "./colors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// "net/http/httputil"
	"os"
	// "bufio"
	// "time"
	"github.com/akamensky/argparse"
	)


// TODO: option for stdout
// var stdout = true
// if outfile is specified, this will become false


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

	// body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Status Code: ", resp.StatusCode)
	// Checks the status of the page
	if resp.StatusCode == 404 {
		// fmt.Println(error("Page doesn't exist"))
	}else{
		// results := re.FindAllString(string(body), -1)
		// fmt.Println("amian")
		// fmt.Println(results)
		// if len(results) == 0 {
		// 	// fmt.Println(error("No proxy was found on "), url)
		// }else {
		// 	fmt.Printf(info("Found %d on %s\n"), len(results), url)
		// 	// write to a file
		// 	for _, proxy := range results {
		// 		// isAlive(proxy)
		// 		// splitProxy := strings.Split(proxy, ":")
		// 		// fmt.Printf("Current proxy: %s, location: %s\n", proxy)
		// 		_, err := outfile.WriteString(proxy + "\n")
		// 		if err != nil {
		// 			fmt.Println("An error occured ", err)
		// 			outfile.Close()
		// 		}
		// 	}
		// }
	}
}

/**
* Parameter: proxyString --> string
* Description: Test if the proxy works by making a request to google.com. if the request code is 200
* 			   then we know it is valid. If otherwise, then we discard it.
**/
func isAlive(proxyString string) bool {
	// Need to set HTTP_PROXY 

	proxyUrl, err := url.Parse("http://"+ proxyString)
	if err != nil {
		panic(err)
	}

	transpot := &http.Transport { Proxy: http.ProxyURL(proxyUrl) }
	client	:= http.Client { Transport: transpot }

	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Alive: ", data)

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

	// Banner - duh
	showBanner()

	// Argument Parsing
	parser := argparse.NewParser("Proxy Scraper", "Proxy Scraper implemented in golang. By PabloPotat0")
	outfile := parser.String("o", "outfile", &argparse.Options{Help: "Good proxies will be stored here"})
	check := parser.Flag("", "check", &argparse.Options{Help: "status of the proxy, alive or dead"})
	chain := parser.String("", "chain", &argparse.Options{Help: "linked list of proxy to direct traffic"})
	// TODO: Only takes in one filter as one
	filter := parser.Selector("f", "filter", []string{"http", "https", "socks"}, &argparse.Options{Help: "Filter the type of proxy you want"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}
	
	// error checking	
	if *outfile == ""{
		fmt.Println(color.ShowWarning("output file not specified. Using default file: good_proxy.txt"))
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


	fmt.Println("Let's begin scraping....")
	if *filter == "" {
		fmt.Println(color.ShowWarning("No filter was specified. Defaulting to http"))
		test.Http_proxies()
	}else if *filter == "http" {
		fmt.Println(color.ShowInfo("Applied filter: http"))
		test.Http_proxies()
	}else if *filter == "https"{
		fmt.Println(color.ShowInfo("Applied filter: https"))
	}else if *filter == "socks"{
		fmt.Println(color.ShowInfo("Applied filter: socks"))
		test.SocksProxies()
	}
	// Create outfile
	// good, _ := os.Create(*outfile)


	// fmt.Println("Let's begin scraping....")
	// test.Http_proxies()
	// scraper("http://globalproxies.blogspot.com/", good)

	// file, err := os.Open(*inputfile)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	//Url validator
	// 	if isValidUrl(scanner.Text()) {
	// 		go scraper(scanner.Text(), good)
	// 		time.Sleep(time.Second)
	// 	}else{
	// 		fmt.Printf(error("%s is not a valid url\n"), scanner.Text())
	// 	}
	// }

	// proxyUrl, err := url.Parse("https://204.19.23.231")
	// if err != nil {
	// 	panic(err)
	// }

	// transpot := &http.Transport { Proxy: http.ProxyURL(proxyUrl) }
	// client	:= http.Client { Transport: transpot }

	// req, err := http.NewRequest("GET", "https://api.ipify.org/", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }

	// data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Alive: ", data)

	// good.Close()
	fmt.Println("Done scraping")
}

