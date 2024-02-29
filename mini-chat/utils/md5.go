package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode lowerCase
func Md5Encode(data string) string {
	encoder := md5.New()
	encoder.Write([]byte(data))
	tempStr := encoder.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MD5Encode upperCase
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword encrypt password
func MakePassword(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}

// ValidatePassword decrypt password
func ValidatePassword(plainPwd, salt, password string) bool {
	return MakePassword(plainPwd, salt) == password
}
