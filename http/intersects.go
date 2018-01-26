package http

import (
	_ "encoding/csv"
	"github.com/thisisaaronland/go-marc/fields"
	_ "github.com/whosonfirst/go-sanitize"
	_ "log"
	gohttp "net/http"
)

func IntersectsHandler() (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
