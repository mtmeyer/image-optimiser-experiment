package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "", "Url to test")
	requests := flag.Int("requests", 0, "Number of requests to execute")
	concurrent := flag.Int("concurrent", 1, "How many concurrent blocks to run")
	flag.Parse()

	if *url == "" || *requests == 0 {
		fmt.Printf("Please pass in both url and requests")
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(*concurrent)

	numberCh := make(chan int, *concurrent * *requests)

	for i := 0; i < *concurrent; i++ {
		go runRequestBlock(*url, *requests, i + 1, wg, numberCh)
	}

  wg.Wait()

}

func runRequestBlock(url string, requests int, currentBlock int, wg *sync.WaitGroup, numberCh chan int) {
	for i := 0; i < requests; i++ {
		time := makeRequest(url)
    numberCh <- currentBlock * (i + 1)
    fmt.Printf("%s \n", time)
	}
  wg.Done()
  
}

func makeRequest(url string) string {
	start := time.Now()

	_, err := http.Get(url)
	if err != nil {
		fmt.Sprintln(os.Stderr, err)
	}

	elapsed := time.Since(start)
	return elapsed.String()
}
