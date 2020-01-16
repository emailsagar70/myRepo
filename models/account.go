package models

//Account is the datamodel of test services
type Account struct {
	Firstname   string  `json:"firstname,omitempty"`
	Middlename  string  `json:"middlename,omitempty"`
	Lastname    string  `json:"lastname,omitempty"`
	Email       string  `json:"email,omitempty"`
	Phonenumber string  `json:"phonenumber,omitempty"`
	Password    string  `json:"password,omitempty"`
	Gender      string  `json:"gender,omitempty"`
	Address     Address `json:"address,omitempty"`
}
