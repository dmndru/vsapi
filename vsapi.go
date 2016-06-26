package vsapi

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://api.vscale.io/v1/"
	userAgent      = "go-vscale/" + libraryVersion
	headerToken    = "X-Token"
	token = "b8042b9503a19b29e3105d0193700573b1f1aa9165b9ffbb2cc416f4ff304b1c"
)

// A Client manages communication with the Vscale API.
type Client struct {
    client *http.Client
    BaseURL *url.URL
}

// Scalet represents a Vscale Scalet
type Scalet struct {
	HostName    string    `json:"hostname,omitempty"`
	Locked      bool      `json:"locked,bool,omitempty"`
	Location    string    `json:"location,omitempty"`
	Rplan       string    `json:"rplan,omitempty"`
	Active      bool      `json:"active,bool,omitempty"`
	Keys		[]Key	  `json:"keys,omitempty"`
	PublicAddress *Addr	  `json:"public_address,omitempty"`
	Status      string    `json:"status,omitempty"`
	MadeFrom	string	  `json:"made_from,omitempty"`
	PrivatAddress *Addr	  `json:"private_address,omitempty"`
	Ctid		int		  `json:"id,float64,omitempty"`
}

// Addr object
type Addr struct {
	Address		string		`json:"address,omitempty"`
	Netmask		string		`json:"netmask,omitempty"`
	Gateway		string		`json:"gateway,omitempty"`
}

// Key object
type Key struct {
	ID		int		`json:"id,float64,omitempty"`
	Name	string	`json:"name,omitempty"`
}

// NewClient creates client for vscale API
func NewClient(httpClient *http.Client) *Client  {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL}
	return c
}
// GetServers returns list of servers
func (c *Client) GetServers() {
	req, err := http.NewRequest("GET", defaultBaseURL + "scalets", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add(headerToken, token)
	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	switch resp.StatusCode {
	case 403:
		fmt.Println("forbidden")
	case 404:
		fmt.Println("path not found")
	case 200:
		scalets, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v+", scalets)
	}
}
