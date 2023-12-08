package test_cases
import (
	"time"
	"io/ioutil"
	"log"
	"fmt"
	"github.com/fatih/color"
	"encoding/json"
	"github.com/tomnomnom/rawhttp"
	"github.com/poseidontor/go-smuggler/internal/structures"
	"github.com/poseidontor/go-smuggler/internal/requests"
)


func Tecl(hostname string, port string, path string, time_out time.Duration)	{
	start := time.Now()
	content, err_io := ioutil.ReadFile("./payloads/tecl.json")
	if err_io != nil	{
		log.Fatal("Error opening file: ", err_io)
	}

	var payload structures.Tecl
	err_file := json.Unmarshal([]byte(content), &payload)
	if err_file != nil {
		log.Fatal("error unmarshaling json: ", err_file)
	}

	var timeout time.Duration = time_out

	var resp[2]*rawhttp.Response
	var err[2]error 
	for i := 0; i < 2; i++	{
		content_type := payload.Cases[i].ContentType
		connection := payload.Cases[i].Connnection
		transfer_encoding := payload.Cases[i].TransferEncoding
		body := payload.Cases[i].Body
		content_length := payload.Cases[i].ContentLength

		httpRequest := structures.HttpRequest{Hostname: hostname, Port: port, 
			Path: path, ContentType: content_type, 
			Connection: connection, TransferEncoding: transfer_encoding, 
			ContentLength: content_length, Body: body, 
			Timeout: timeout}

		resp[i], err[i] = requests.MakeRequestHttp(httpRequest);
	}
	if(err[0] == nil && err[1] == nil){
		if resp[0] != nil && resp[1] != nil{ 
			if resp[0].StatusCode() == "504" || resp[1].StatusCode() == "504" {
				fmt.Printf("[-] Received 504: Gateway Timeout for %s\n", hostname)
			}

			if resp[0].StatusCode() == "500" && resp[1].StatusCode() == "200" {
				color.Red("[+] https://%s:%s%s could possibly be vulnerble to TE.CL based request smuggling.\n", hostname, port, path)
			}
			secs := time.Since(start).Seconds()

			fmt.Printf("%.2f seconds taken for https://%s:%s%s\n", secs, hostname, port, path)
		}
	} else {
		fmt.Printf("Request for %s errored out!\n" , hostname)
	}
}