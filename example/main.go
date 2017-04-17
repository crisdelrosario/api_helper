package main

import (
	"github.com/crisdelrosario/api_helper"
)

func main() {
	testApi := api.New("http://localhost/fake/api", api.ContentTypeJSON)
	if _, err := testApi.Get("hello", nil); err != nil {
		panic(err.Error())
	} else {
	}
}
