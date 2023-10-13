package structures

type Clte struct	{
	Cases []struct	{
		ContentType string `json:"content_type"`
		Connnection string `json:"connection"`
		TransferEncoding string `json:"transfer_encoding"`
		Body string `json:"body"`
	} `json:"cases"`
}