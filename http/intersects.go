package http

import (
       _ "github.com/thisisaaronland/go-marc/fields"
	_ "github.com/whosonfirst/go-sanitize"
	"log"
	gohttp "net/http"
)

func IntersectsHandler() (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		fh, _, err := req.FormFile("upload")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		log.Println(fh)

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
