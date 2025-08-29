package util

import (
	"encoding/base64"
)

func GenAuthToken(account string, password string) string {
	// BasicAuth
	return base64.StdEncoding.EncodeToString([]byte(account + ":" + password))
}
