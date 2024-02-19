package myPkg

const BaseURL="https://api.shodan.io/"
// if shodan changes the link 


type Client struct{

	apiKey string
}


// Creating an instances of Client

func New(apiKey string) *Client{
	return &Client{apiKey: apiKey}
}