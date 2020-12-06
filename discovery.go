package main

import "fmt"

var (
	ApiRoots = []string{
		"/api1/",
		"/trustgroup1/",
		"/trustgroup2/",
	}
)

type DiscoveryResource struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	Contact string `json:"contact,omitempty"`
	Default string `json:"default,omitempty"`
	ApiRoots []string `json:"api_roots,omitempty"`
}

func ProcessDiscovery() DiscoveryResource {
	dr := DiscoveryResource{}
	dr.Description = "TAXII 2.1 Server in GoLang"
	dr.Title = "TAXII 2.1 Server"
	dr.Contact = "Kiet T. Tran, Ph.D."
	dr.Default = fmt.Sprintf("https://%s%s",HostName, ApiRoots[0])
	dr.ApiRoots = make([]string, 0, len(ApiRoots))
	for _, apiRoot := range ApiRoots {
		dr.ApiRoots = append(dr.ApiRoots, fmt.Sprintf("https://%s%s",HostName, apiRoot))
	}
	return dr
}
