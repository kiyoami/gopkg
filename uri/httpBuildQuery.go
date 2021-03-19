package uri

import (
	"fmt"
	"net/url"
	"strings"
)

// HttpBuildQuery : Map converts to uri query string.
func HttpBuildQuery(params map[string]interface{}) string {
	var x []string
	for key, value := range params {
		x = append(x, fmt.Sprintf("%s=%s", key, url.QueryEscape(fmt.Sprint(value))))
	}
	return strings.Join(x, "&")
}
