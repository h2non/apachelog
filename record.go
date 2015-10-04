package apachelog

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const Pattern = "%s - - [%s] \"%s\" %d %d %.4f\n"

type Record struct {
	http.ResponseWriter
	status                int
	responseBytes         int64
	ip                    string
	method, uri, protocol string
	time                  time.Time
	elapsedTime           time.Duration
}

func (r *Record) Log(out io.Writer) {
	timeFormat := r.time.Format("01/Jan/2006 03:04:05")
	request := fmt.Sprintf("%s %s %s", r.method, r.uri, r.protocol)
	fmt.Fprintf(out, Pattern, r.ip, timeFormat, request, r.status, r.responseBytes, r.elapsedTime.Seconds())
}

func (r *Record) Write(p []byte) (int, error) {
	written, err := r.ResponseWriter.Write(p)
	r.responseBytes += int64(written)
	return written, err
}

func (r *Record) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}
