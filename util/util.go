package util

import (
	"encoding/json"
	pkgErrors "github.com/pkg/errors"
)

func MustMarshalToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(pkgErrors.WithStack(err))
	}
	return string(b)
}
