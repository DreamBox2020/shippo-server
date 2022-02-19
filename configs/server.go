package configs

type Server struct {
	Addr         string `json:"addr"`
	CookieDomain string `json:"cookieDomain"`
	UploadDir    string `json:"uploadDir"`
}
