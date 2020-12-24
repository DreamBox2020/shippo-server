package configs

type Sms struct {
	RegionId        string `json:"regionId"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SignName        string `json:"signName"`
	TemplateCode    string `json:"templateCode"`
	TestPhoneNumber string `json:"testPhoneNumber"`
}
