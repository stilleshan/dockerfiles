package ga

import (
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// PageView sends analysis data of t=pageview
func PageView(data CommonData) error {
	data.HitType = "pageview"

	v, err := query.Values(data)
	if err != nil {
		return errors.Wrap(err, "could not encode query")
	}

	err = send(v.Encode())

	return nil
}

// Timing sends analysis data of t=timing
func Timing(data CommonData, tData TimingData) error {
	data.HitType = "timing"

	v1, err := query.Values(data)
	if err != nil {
		return errors.Wrap(err, "could not encode query")
	}
	v2, err := query.Values(tData)
	if err != nil {
		return errors.Wrap(err, "could not encode query")
	}
	concatURLValues(v1, v2)

	err = send(v1.Encode())

	return nil
}

// Detect tests network connection
func Detect() error {
	err := send("")
	return err
}
