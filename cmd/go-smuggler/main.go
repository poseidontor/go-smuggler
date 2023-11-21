package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"flag"
	"strings"
	"time"
	"os"
	"sync"
	"net/url"
	"github.com/poseidontor/go-smuggler/internal/test_cases"
)


func main()	{
	
	var uri string
	var file_path string
	var time_out time.Duration
	flag.StringVar(&uri, "u", "", "please provide the URL")
	flag.DurationVar(&time_out, "t", 10, "timeout for requests")
	flag.StringVar(&file_path, "f", "", "file containing multiple urls")
	flag.Parse()

	urls := make([]string, 0)

	if uri != ""	{
		urls = append(urls, uri)
	}

	if file_path != ""	{
		url_file, err_io := ioutil.ReadFile(file_path)
		if err_io != nil	{
			log.Fatal("Error opening file. Please provide the full path to the file.")
		}
		read_urls := strings.Split(string(url_file), "\n")
		for _,link	:= range read_urls	{
			urls = append(urls, link)
		}
	}

	if uri == "" && file_path == ""	{
		fmt.Println("Usage: go run cmd/main.go -u {URL} [-f {urls_file}] [-t {timeout}]")
		os.Exit(0)		
	}

	
	var wg sync.WaitGroup
   	wg.Add(1)
	go scan(urls, time_out, &wg)
	//time.Sleep(5 * time.Second)
	wg.Wait()
	
}

func scan(urls []string, time_out time.Duration, wg *sync.WaitGroup)	{

	defer wg.Done()

	var hostname string
	var port string
	var path string
	var u *url.URL 
	var err error
	
	
	for _, uri := range urls	{
		if uri != ""	{
			u, err = url.Parse(uri)
			if err != nil {
				log.Fatal("Error reading the URL")
			}
			hostname = u.Hostname()
			port = u.Port()
			path = u.Path

			if port == ""	{
				port = "443"
			}	

			test_cases.Clte(hostname, port, path, time_out)
			test_cases.Tecl(hostname, port, path, time_out)
			test_cases.Tete(hostname, port, path, time_out)
		}
	}
	
}
