package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func main() {
	params := map[string]string{
		"app_id":     "1000000000[PHONE]",
		"charset":    "UTF-8",
		"format":     "json",
		"msg_id":     "5eb0ed071af5434da4cc1942ac42c174",
		"notify_url": "<https://www.scgsj.com/notify.do",
		"sign_type":  "RSA",
		"timestamp":  "2017-07-13 19:48:47",
		// other request parameters
		"biz_content": "{\"authen_acct_name\":\"x9G3wg==\",\"authen_acct_no\":\"[SSN][PHONE]\",\"authen_name\":\"x9G3wg==\",\"auto_turn_flag\":\"0\",\"cert_no\":\"42876719[PHONE]\",\"cert_type\":\"0\",\"language\":\"ZH_CN\",\"logon_id\":\"02000[PHONE].p.0200\",\"notify_type\":\"HS\",\"request_ip\":\"192.168.1.1\",\"tran_time\":\"2016[PHONE]\",\"verified_corp_id\":\"2000EG0000136\",\"verified_corp_name\":\"uaTJzL7W\",\"verified_flag\":\"1\",\"verified_id\":\"800136\",\"verified_info\":\"1eLKx9K7uPbHqcP7xNrI3cW2\",\"verified_kind\":\"0\",\"verified_type\":\"0\"}",
	}

	// Step 1: Filter parameters and sort
	var keys []string
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// Step 2: Construct the string to sign
	var signStr strings.Builder
	for _, k := range keys {
		signStr.WriteString(fmt.Sprintf("%s=%s&", k, params[k]))
	}
	signStrString := signStr.String()
	signStrString = signStrString[:len(signStrString)-1] // Remove the trailing '&'

	// Step 3: Load private key
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048) // Replace with your own private key

	// Step 4: Sign the string
	h := sha1.New()
	h.Write([]byte(signStrString))
	hashed := h.Sum(nil)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed)

	// Step 5: Base64 encode the signature
	base64Signature := base64.StdEncoding.EncodeToString(signature)
	fmt.Println(base64Signature)

	formData := url.Values{}
	for key, value := range params {
		formData.Set(key, value)
	}
	formData.Set("sign", base64Signature)

	// Step 7: Create the HTTP request
	req, _ := http.NewRequest("POST", "your_api_url_here", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "your_gateway_domain_here")

	// Step 8: Send the HTTP request
	client := &http.Client{Timeout: time.Second * 10} // Replace with your own HTTP client configuration
	resp, _ := client.Do(req)

	fmt.Println(resp)

	// Step 9: Process the response
	// ...
}
