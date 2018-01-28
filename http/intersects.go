package http

import (
	"github.com/thisisaaronland/go-marc/fields"
	"github.com/whosonfirst/go-sanitize"
	"github.com/whosonfirst/go-whosonfirst-csv"
	"io"
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

		// sudo put me in a function... (20180128/thisisaaronland)

		reader, err := csv.NewDictReader(fh)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		sanitize_opts := sanitize.DefaultOptions()

		for {
			row, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			marc_raw, ok := row["marc_034"]

			if !ok {
				continue		// throw an error?
			}

			marc_clean, err := sanitize.SanitizeString(marc_raw, sanitize_opts)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			// log.Println("RAW", marc_raw)
			// log.Println("CLEAN", marc_clean)

			bounds, err := fields.Parse034(marc_clean)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			bbox, err := bounds.BoundingBox()

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}

			log.Println(bbox)

			// FIND INTERSECTING THINGS HERE...
		}

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
