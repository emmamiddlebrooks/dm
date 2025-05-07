package internal

import "strings"

func GetSubdir(host string) string {
	if strings.HasPrefix(host, "wildleap.") {
		return "wildleap"
	}
	return "main"
}
