package models

//Account is the datamodel of test services
type Account struct {
	ID          string `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Middlename  string `json:"middlename,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Email       string `json:"email,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty"`
	Password    string `json:"password,omitempty"`
	Gender      string `json:"gender,omitempty"`
	// here pointer is used to avoid display empty struct in UI, if bad data is received
	Address     *Address `json:"address,omitempty"`
	Status      string   `json:"status,omitempty"`
	OldPassword string   `json:"oldpassword,omitempty"`
}
