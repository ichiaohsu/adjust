package main

import (
	"flag"
	"fmt"
	"sync"
)

type errorPair [2]string

func main() {
	// setup parallels client
	var parallels int
	flag.IntVar(&parallels, "parallel", 10, "parallel client amount")
	flag.Parse()

	urls, err := parseURL(flag.Args())
	if err != nil {
		panic(err)
	}

	data := make(chan string, parallels)
	// single error collector
	errorChan := make(chan errorPair)
	errorMap := make(map[string]string)
	go func(errChan chan errorPair, errMap map[string]string) {
		for err := range errChan {
			if _, ok := errMap[err[0]]; !ok {
				errMap[err[0]] = err[1]
			}
		}
	}(errorChan, errorMap)

	go func() {
		for _, u := range urls {
			data <- u
		}
		close(data)
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < parallels; i++ {
		wg.Add(1)
		go func() {
			c := newSimpleClient()
			for url := range data {
				response, err := c.getHash(url)
				if err != nil {
					errorChan <- errorPair{url, err.Error()}
					continue
				}
				fmt.Printf("%s %s\n", url, response)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	// close error channel,
	// print out error results after all other successful results
	close(errorChan)
	if len(errorMap) != 0 {
		fmt.Println("-----")
		for url, e := range errorMap {
			fmt.Printf("%s error:%s\n", url, e)
		}
	}
}
