package types

type Address struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CompanyName string `json:"companyName"`
	Address     string `json:"address"`
	City        string `json:"city"`
	County      string `json:"county"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Phone1      string `json:"phone1"`
	Phone2      string `json:"phone2"`
	Email       string `json:"email"`
	Web         string `json:"web"`
	Uuid        string `json:"uuid"`
}
