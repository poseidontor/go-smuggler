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


func Clte(hostname string, port string, path string, time_out time.Duration)	{
	start := time.Now()
	content, err_io := ioutil.ReadFile("./payloads/clte.json")
	if err_io != nil	{
		log.Fatal("Error opening file: ", err_io)
	}

	var payload structures.Clte
	err_file := json.Unmarshal([]byte(content), &payload)
	if err_file != nil {
		log.Fatal("error unmarshaling json: ", err_file)
	}

	var timeout time.Duration = time_out

	var resp[2]*rawhttp.Response
	for i := 0; i < 2; i++	{
		content_type := payload.Cases[i].ContentType
		connection := payload.Cases[i].Connnection
		transfer_encoding := payload.Cases[i].TransferEncoding
		body := payload.Cases[i].Body

		resp[i] = requests.Make_request_http1(hostname, port, path, 
		content_type, connection, transfer_encoding, "", body, timeout);
	}
	
	if resp[0].StatusCode() == "504" || resp[1].StatusCode() == "504" {
		fmt.Printf("[-] Received 504: Gateway Timeout for %s\n", hostname)
	}

	if resp[0].StatusCode() == "500" && resp[1].StatusCode() == "200" {
		color.Red("[+] https://%s:%s%s could possibly be vulnerble to CL.TE based request smuggling.\n", hostname, 
	port, path)
	}
	secs := time.Since(start).Seconds()

	fmt.Printf("%.2f seconds taken for https://%s:%s%s\n", secs, hostname, port, path)
}