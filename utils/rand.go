package utils

import (
	"math/rand"
	"time"
)

const uidLen = 10
const accessTokenLen = 20
const uidPrefix = "u_"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetNumber = "0123456789"

func genRandomString(prefix string, strLen int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, strLen)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return prefix + string(b)
}

func GenerateAccessToken() string {
	return genRandomString("", accessTokenLen, charset)
}
func GenerateUserID() string {
	return genRandomString(uidPrefix, uidLen, charset)
}
func GenerateVerifyCode() string {
	return genRandomString("", 6, charsetNumber)
}
func GenerateAppID() string {
	return genRandomString("app_", 10, charset)
}
func GenerateTaskID() string {
	return genRandomString("", 10, charset)
}
func GenerateSKU() string {
	return genRandomString("sku_", 10, charset)
}
func GenerateOrderID() string {
	return genRandomString("odr_", 10, charset)
}

func GenStringWithPrefix(prefix string, strLen int) string {
	return genRandomString(prefix, strLen, charset)
}
