package structures

import (
	"time"
)

type HttpRequest struct	{
	
	Hostname string
	Port string
	Path string
	ContentType string
	Connection string
	TransferEncoding string
	ContentLength string
	Body string
	Timeout time.Duration
	
}
