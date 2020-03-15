package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"flag"
	"os"
	"bufio"
	"sync"
	"time"
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

func main(){
	// go run main.go -o outfile.txt -i inputfile.txt 
	// Need more options
	// --check: status of the proxy, alive or dead
	// --chain: linked list of proxy to direct traffic

	// 1) read file from users -- DONE
	// 2) For each line in the file(verify if it is a url)
	// 2A) CHECK IF URL IS "ALIVE" -- Sort of
	// 3) Create a thread for each line -- DONE
	// 4) then run the scraper -- DONE

	// Argument Parsing
	inputfile := flag.String("i", "proxylist.txt", "a list of proxy in a text file")
	outfile := flag.String("o", "good_proxy.txt", "Results will be stored here")
	flag.Parse()

	// Create outfile
	good, _ := os.Create(*outfile)

	// Banner - duh
	showBanner()

	fmt.Println("input: ", *inputfile)
	fmt.Println("output: ", *outfile)
	fmt.Println("Let's begin scraping....")

	file, err := os.Open(*inputfile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go scraper(scanner.Text(), good)
		time.Sleep(time.Second)
	}

	good.Close()
	fmt.Println("Ending")
}

