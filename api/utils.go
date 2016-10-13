package api

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Record struct {
	Url    string
	Method string
	Header map[string][]string
	Body   string
}

// Decode request body by content encoding
func DecodeRequestBody(r *http.Request) ([]byte, error) {

	ce, _ := r.Header["Content-Encoding"]

	for _, x := range ce {
		switch x {
		case "gzip":
			gr, e := gzip.NewReader(r.Body)
			defer gr.Close()
			if e != nil {
				return nil, e
			}
			return readAll(gr)
		case "deflate":
			fr, e := zlib.NewReader(r.Body)
			defer fr.Close()
			if e != nil {
				return nil, e
			}
			return readAll(fr)
		}
	}

	// No Content-Encoding header detected
	return readAll(r.Body)
}

// Read from reader to string
func readAll(r io.Reader) ([]byte, error) {
	if body, err := ioutil.ReadAll(r); err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func ParseStringTag(sa []string) map[string]string {
	tags := make(map[string]string, 0)
	for _, s := range sa {
		ss := strings.Split(s, ":")
		if len(ss) == 2 {
			tags[ss[0]] = ss[1]
		}
	}
	return tags
}
