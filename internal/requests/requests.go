package requests

import (
	"fmt"
	"log"
	"github.com/tomnomnom/rawhttp"
	"github.com/poseidontor/go-smuggler/internal/structures"
	"time"
)

func MakeRequestHttp(httpRequest structures.HttpRequest) (resp *rawhttp.Response) {
	
	req, err_http := rawhttp.FromURL("POST", "https://"+httpRequest.Hostname+"/")
	if err_http != nil {
		log.Fatal(err_http)
	}
		
	req.Method = "POST"
	req.Proto = "HTTP/1.1"
	req.Hostname = httpRequest.Hostname
	req.Port = httpRequest.Port
	req.Path = httpRequest.Path

	//automatically assign host header value
	req.AutoSetHost()

	//setting headers
	req.AddHeader(httpRequest.ContentType)
	req.AddHeader(httpRequest.Connection)
	req.AddHeader(httpRequest.TransferEncoding)

	req.Body = httpRequest.Body
	req.Timeout = httpRequest.Timeout * time.Second
	
	//Automatically set content-length if not specified
	if httpRequest.ContentLength == ""	{
		req.AutoSetContentLength()
	}	else {
		req.AddHeader(httpRequest.ContentLength)
	}
	
	resp, err_req := rawhttp.Do(req)
	if err_req != nil {
		fmt.Printf("[-] Connection Error: %s\n", err_req)
	}
	return
}