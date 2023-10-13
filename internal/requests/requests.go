package requests

import (
	"fmt"
	"log"
	"github.com/tomnomnom/rawhttp"
	"time"
)

func Make_request_http1(hostname string, port string, path string, content_type string, 
	connection string, transfer_encoding string, content_length string, body string, timeout time.Duration ) (resp *rawhttp.Response) {
	
	req, err_http := rawhttp.FromURL("POST", "https://"+hostname+"/")
	if err_http != nil {
		log.Fatal(err_http)
	}
		
	req.Method = "POST"
	req.Proto = "HTTP/1.1"
	req.Hostname = hostname
	req.Port = port
	req.Path = path

	//automatically assign host header value
	req.AutoSetHost()

	//setting headers
	req.AddHeader(content_type)
	req.AddHeader(connection)
	req.AddHeader(transfer_encoding)

	req.Body = body
	req.Timeout = timeout * time.Second
	
	//Automatically set content-length if not specified
	if content_length == ""	{
		req.AutoSetContentLength()
	}	else {
		req.AddHeader(content_length)
	}
	//fmt.Printf("%s\n\n", req.String())
	
	resp, err_req := rawhttp.Do(req)
	if err_req != nil {
		fmt.Printf("[-] Connection Error: %s\n", err_req)
	}
	return
}