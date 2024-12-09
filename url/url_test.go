package url

import (
	"fmt"
	"net/url"
	"testing"
)

func TestX(t *testing.T) {
	s := url.PathEscape("http://wdt.wangdian.cn/openapi?2023-09-13 17:00:15")
	fmt.Println(s)
	s = url.QueryEscape(s)
	fmt.Println(s)

	parsedURL, _ := url.ParseRequestURI("https://domain/")
	parsedURL.RawQuery = "v=1.0&end_time=2023-09-13 17:00:15"
	// urlValues := parsedURL.Query()
	// urlValues.Add("v", "1.0")
	// urlValues.Add("end_time", "2023-09-13 17:00:15")
	fmt.Println(parsedURL.String())
	fmt.Println(parsedURL.EscapedPath())
}
