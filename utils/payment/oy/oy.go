package oy

type configOy struct {
	Username string `default:""`
	ApiKey   string `default:""`
	Endpoint string `default:""`
}

var oyConfig configOy

func SetupOyPayment(config configOy) {
	oyConfig = config
}

func GetOyConfig() configOy {
	return oyConfig
}

func CreatePayment(payment string) (string, error) {
	return payment, nil
}
