package uri

import "testing"

func TestHttpBuildQuery(t *testing.T) {
	params := map[string]interface{}{
		"id":   123,
		"name": "Oami Kiyokazu",
	}
	query := HttpBuildQuery(params)
	if query != "id=123&name=Oami+Kiyokazu" {
		t.Fatal("convert error")
	}
}
