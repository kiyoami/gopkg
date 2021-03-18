package tree

import (
	"log"
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	log.SetFlags(log.Llongfile)
	var obj Tree

	if err := obj.LoadJson([]byte(invalidJson)); err != nil {
		log.Println("Expected error:", err)
	} else {
		t.Error("LoadJson() error.", err)
	}

	if err := obj.LoadJson([]byte(validJson)); err != nil {
		t.Error("LoadJson() error.", err)
	}

	if v, err := obj.Get("fff.eee.1"); err != nil {
		t.Error("Get() error.", err)
	} else if reflect.TypeOf(v).Kind() != reflect.Float64 {
		t.Errorf("type error. %v (%T)", v, v)
	}

	if _, err := obj.Get("zzz.eee.1"); err != nil {
		log.Println("Expected error:", err)
	} else {
		t.Error("Get() error.")
	}

	if _, err := obj.Get("fff.eee.5"); err != nil {
		log.Println("Expected error:", err)
	} else {
		t.Error("Get() error.")
	}

	if _, err := obj.Get("fff.eee.aaa"); err != nil {
		log.Println("Expected error:", err)
	} else {
		t.Error("Get() error.")
	}
}

const invalidJson = `{aaa}`

const validJson = `
{
	"aaa": 1,
	"bbb": "BBB",
	"ccc": true,
	"ddd": null,
	"eee": [1,2,3],
	"fff": {
		"aaa": 1,
		"bbb": "BBB",
		"ccc": true,
		"ddd": null,
		"eee": [1,2,3]
	}
}
`
