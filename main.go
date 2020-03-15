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
	)

var red = "\033[1;31m"
var green = "\033[1;32m"
var yellow = "\033[1;33m"
var defcol = "\033[0m"

var re = regexp.MustCompile(`[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}(\:[\d]{2,5})?`)

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
func info(msg string) string{
	return red + "[" + green + "+" + red + "] - " + defcol + msg
}

func error(msg string) string {
	return red + "[" + yellow + "!" + red + "] - " + defcol + msg
}

func scraper(url string, outfile *os.File) {
	// Split result to good proxy and bad proxy
	resp, err := http.Get(url) 

	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Status Code: ", resp.StatusCode)
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
				_, err := outfile.WriteString(proxy + "\n")
				if err != nil {
					fmt.Println("An error occured ", err)
					outfile.Close()
				}
			}
		}
	}
}

func isAlive(url string) bool {
	return true
}

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
	// showBanner()

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

