package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"spiderden.org/masta"
	"github.com/urfave/cli/v2"
)

func testWithServer(h http.HandlerFunc, testFuncs ...func(*cli.App)) string {
	ts := httptest.NewServer(h)
	defer ts.Close()

	cli.OsExiter = func(n int) {}

	client := masta.NewClient(&masta.Config{
		Server:       ts.URL,
		ClientID:     "foo",
		ClientSecret: "bar",
		AccessToken:  "zoo",
	})

	var buf bytes.Buffer
	app := makeApp()
	app.Writer = &buf
	app.Metadata = map[string]interface{}{
		"client": client,
		"config": &masta.Config{
			Server: "https://example.com",
		},
		"xsearch_url": ts.URL,
	}

	for _, f := range testFuncs {
		f(app)
	}

	return buf.String()
}
