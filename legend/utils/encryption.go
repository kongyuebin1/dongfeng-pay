package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"strings"
)

//type Encrypt struct{}

func EncodeMd5(str string) string {
	hash := md5.New()
	buf := []byte(str)
	hash.Write(buf)

	return hex.EncodeToString(hash.Sum(nil))
}

//base64编码
func Base64Encode(raw []byte) string {
	t := base64.StdEncoding.EncodeToString(raw)
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "\r", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "\n\r", "", -1)
	t = strings.Replace(t, "\r\n", "", -1)
	return t
}

func getMd5() string {
	return ""
}

//base64解码
func Base64Decode(raw string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(raw)
}

//SHA1withRSA签名
func SHA1withRSA(context string) string {

	_, block := pem.Decode([]byte(`-----BEGIN RSA PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAIYmcOHeAztObFnQX+qIf/e7/7BtFtYaaijNN1P59GQmPm0R8XE4s8afvX/r5eBjX9muJ51x/8PaXzw05vjoHDhJnGLeABtuem3PNgg+vfk09HrA4yhDfqU6GJPKOWF6ko3DprEDFUijC6lKJFOhsF7vQu8ZnYsCeuBpQmetfG1ZAgMBAAECgYBIPhlaOXIqFQCamXGd3uZzJgX7H7RFlrIGyQT7r0biTAogOKJ6Y5vE4i9t3T7NSRbMJlJlIognE8lnpeGgt3bCPYisxsYd6Eo9YA6stzNmE9Pwy5gPi6F+qhBFkbbjETP0XF/PzRLZnRGKMPAQHufbA/qpQYsWsrGAn7wKCnL+0QJBAL8kKw2tdpThcFzEOHjvCCkcDd0wUf8mqXz+TqNgzvCKfUv+PIyfETF6Gd+mrh4o2PLqBM53xGv45nraXFL3n+UCQQCzq6LjOfSDTuPatPlYAMsWNA/1zfUApzRijLKfY7+Ek6BEVJ/NuYR0jlLAP3JiNsUhEw/LAO0GsJtuHXcAC3hlAkEAgIyml/BNjBuCIiGliU/ZQSyo9lWFEADEhFfUM3TsOEIrumwl9L0WJxxjQlMrTwVRwy04RlOuOp+PApjQ9surMQJAQAzIxZ5Md17xRW9MkD3AKEspAWSJmdEBkLw9lSqXBKkn8hQE3+7ptC9keppjqXWC8tZ7w8+xr7fXwPqKCJ8OLQJBAJmxZnpUp5lKYiHyiNGXrPukleyUJ4BoLiF0/gvt+TTg1YQ/nkK+V/3YurAabyUh878hasRzjSW9ZL8B94aj4kc=
		-----END RSA PRIVATE KEY-----`))
	if block == nil {
		panic("私钥错误")
	}

	private, err := x509.ParsePKCS8PrivateKey(block) //之前看java demo中使用的是pkcs8
	if err != nil {
		return ""
	}
	h := crypto.Hash.New(crypto.SHA1) //进行SHA1的散列
	h.Write([]byte(context))
	hashed := h.Sum(nil)

	// 进行rsa加密签名
	signedData, err := rsa.SignPKCS1v15(rand.Reader, private.(*rsa.PrivateKey), crypto.SHA1, hashed)

	data := base64.StdEncoding.EncodeToString(signedData)
	return data
}
