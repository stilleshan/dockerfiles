package ga

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

const gaURL = "https://www.google-analytics.com/collect"

var httpClient = &http.Client{}

func send(qs string) error {
	req, err := http.NewRequest(http.MethodPost, gaURL, strings.NewReader(qs))
	if err != nil {
		return errors.Wrap(err, "could not create request")
	}

	// https://golang.org/pkg/net/http/#Client.Do
	// On error, any Response can be ignored. A non-nil Response with a non-nil error only occurs when
	// CheckRedirect fails, and even then the returned Response.Body is already closed.
	resp, err := httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "could not make request")
	}
	defer resp.Body.Close()

	return nil
}

func concatURLValues(v1 url.Values, v2 url.Values) {
	for key, values := range v2 {
		if len(v2.Get(key)) != 0 && len(v1.Get(key)) == 0 {
			// do not replace existed key
			v1[key] = values
		}
	}
}
