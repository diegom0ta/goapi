package apiParser

import (
	"bytes"
	"io"
)

func WriteResponse(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	respBytes := buf.String()

	return string(respBytes)
}
