package response

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Success  bool        `json:"success"`
	Session  string      `json:"session"`
	Resource interface{} `json:"resource"`
	Sign     string      `json:"sign"`
	Other    interface{} `json:"other"`
}
