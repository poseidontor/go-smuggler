
package structures

type Tete struct	{
	Cases []struct	{
		ContentType string `json:"content_type"`
		Connnection string `json:"connection"`
		TransferEncoding string `json:"transfer_encoding"`
		ContentLength string `json:"content_length"`
		Body string `json:"body"`
	} `json:"cases"`
}