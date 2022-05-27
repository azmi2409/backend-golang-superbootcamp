package duitku

import (
	"api-store/utils"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Payment struct {
	MerchantId     string `json:"merchant_id"`
	Amount         string `json:"amount"`
	OrderId        string `json:"order_id"`
	Details        string `json:"details"`
	Email          string `json:"email"`
	PaymentMethod  string `json:"payment_method"`
	CusVAName      string `json:"cus_va_name"`
	PhoneNumber    string `json:"phone_number"`
	ItemDetails    string `json:"item_details"`
	CustomerDetail string `json:"customer_detail"`
	ReturnURL      string `json:"return_url"`
	CallbackURL    string `json:"callback_url"`
	Signature      string `json:"signature"`
	ExpiryPeriod   string `json:"expiry_period" default:"1200"`
	Shopee         string `json:"shopee"`
}

var config Payment

var DuitkuURL = utils.GetEnv("DUITKU_URL", "https://api.duitku.com/v1/payment/create")
var APIKey = utils.GetEnv("DUITKU_API_KEY", "")

func GenerateMD5Signature(merchantCode, merchantOrderID, amount, apiKey string) string {
	join := merchantCode + merchantOrderID + amount + apiKey
	byteData := []byte(join)
	hash := md5.Sum(byteData)

	output := fmt.Sprintf("%x", hash)
	return output
}

func CreateTransaction(payment Payment) (string, error) {
	payment.Signature = GenerateMD5Signature(payment.MerchantId, payment.OrderId, payment.Amount, APIKey)
	client := &http.Client{}
	req, err := http.NewRequest("POST", DuitkuURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+APIKey)
	req.Header.Add("X-API-KEY", APIKey)
	req.Header.Add("X-MERCHANT-CODE", payment.MerchantId)
	req.Header.Add("X-MERCHANT-ORDER-ID", payment.OrderId)

	body, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func GetPaymentMethod(p Payment) []string {

	url := "https://sandbox.duitku.com/webapi/api/merchant/paymentmethod/getpaymentmethod"

	params := []string{
		p.MerchantId,
		p.Amount,
		time.Now().Format("yyyy-MM-dd HH:mm:ss"),
		GenerateMD5Signature(p.MerchantId, p.OrderId, p.Amount, APIKey),
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	body, _ := json.Marshal(params)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return []string{string(res)}
}
