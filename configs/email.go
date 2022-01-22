package configs

type Email struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	AddressName string `json:"addressName"`
	TestEmail   string `json:"testEmail"`
}
