
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


func Tete(hostname string, port string, path string, time_out time.Duration)	{
	start := time.Now()
	content, err_io := ioutil.ReadFile("./payloads/tete.json")
	if err_io != nil	{
		log.Fatal("Error opening file: ", err_io)
	}

	var payload structures.Tete
	err_file := json.Unmarshal([]byte(content), &payload)
	if err_file != nil {
		log.Fatal("error unmarshaling json: ", err_file)
	}

	var timeout time.Duration = time_out
	
	var resp[2]*rawhttp.Response
	
	for _, value := range payload.Cases {
		for i := 0; i < 2; i++	{
			content_type := value.ContentType
			connection := value.Connnection
			transfer_encoding := value.TransferEncoding
			body := value.Body
			content_length := value.ContentLength

			resp[i] = requests.Make_request_http1(hostname, port, path, content_type, connection, transfer_encoding, content_length, body, timeout);
			time.Sleep(10 * time.Millisecond)
		}
		
		if resp[0] != nil && resp[1] != nil{
			if resp[0].StatusCode() == "504" || resp[1].StatusCode() == "504" {
				fmt.Printf("[-] Received 504: Gateway Timeout for %s\n", hostname)
			}

			if resp[0].StatusCode() == "200" && resp[1].StatusCode() == "403" {
				color.Red("[+] https://%s:%s%s could possibly be vulnerble to TE.TE based request smuggling.\n", hostname, port, path)
				break
			}
			secs := time.Since(start).Seconds()

			fmt.Printf("%.2f seconds taken for https://%s:%s%s\n", secs, hostname, port, path)
		}
		
	}
} 
	
