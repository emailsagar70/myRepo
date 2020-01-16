package models

// Address Represents the Address
type Address struct {
	Housenumber string `json:"housenumber,omitempty"`
	Streetname  string `json:"streetname,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Zipcode     string `json:"zipcode,omitempty"`
}
