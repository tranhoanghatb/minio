/*
 * Minio Cloud Storage, (C) 2015, 2016, 2017 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// recordRequest - records the first recLen bytes
// of a given io.Reader
type recordRequest struct {
	// Data source to record
	io.Reader
	// Response body should be logged
	logBody bool
	// Internal recording buffer
	buf bytes.Buffer
}

func (r *recordRequest) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	if r.logBody {
		r.buf.Write(p[:n])
	}
	if err != nil {
		return n, err
	}
	return n, err
}

// Return the bytes that were recorded.
func (r *recordRequest) Data() []byte {
	return r.buf.Bytes()
}

// recordResponseWriter - records the first recLen bytes
// of a given http.ResponseWriter
type recordResponseWriter struct {
	// Data source to record
	http.ResponseWriter
	// Response body should be logged
	logBody bool
	// Internal recording buffer
	headers bytes.Buffer
	body    bytes.Buffer
	// The status code of the current HTTP request
	statusCode int
	// Indicate if headers are written in the log
	headersLogged bool
}

// Write the headers into the given buffer
func writeHeaders(w io.Writer, statusCode int, headers http.Header) {
	fmt.Fprintf(w, "%d %s\n", statusCode, http.StatusText(statusCode))
	for k, v := range headers {
		fmt.Fprintf(w, "%s: %s\n", k, v[0])
	}
}

func (r *recordResponseWriter) WriteHeader(i int) {
	r.statusCode = i
	if !r.headersLogged {
		writeHeaders(&r.headers, i, r.ResponseWriter.Header())
		r.headersLogged = true
	}
	r.ResponseWriter.WriteHeader(i)
}

func (r *recordResponseWriter) Write(p []byte) (n int, err error) {
	n, err = r.ResponseWriter.Write(p)
	// We log after calling ResponseWriter.Write() because this latter
	// proactively prepares headers when WriteHeader is not called yet.
	if !r.headersLogged {
		writeHeaders(&r.headers, http.StatusOK, r.ResponseWriter.Header())
		r.headersLogged = true
	}
	if r.statusCode != http.StatusOK && r.statusCode != http.StatusNoContent && r.statusCode != 0 {
		// We always log error responses.
		r.body.Write(p)
		return
	}
	if r.logBody {
		r.body.Write(p)
		return
	}
	return n, err
}

func (r *recordResponseWriter) Flush() {
	r.ResponseWriter.(http.Flusher).Flush()
}

// Return response headers.
func (r *recordResponseWriter) Headers() []byte {
	return r.headers.Bytes()
}

// Return response body.
func (r *recordResponseWriter) Body() []byte {
	return r.body.Bytes()
}

// Log headers and body.
func httpTracelogAll(f http.HandlerFunc) http.HandlerFunc {
	if !globalHTTPTrace {
		return f
	}
	return httpTracelog(f, true)
}

// Log only the headers.
func httpTracelogHeaders(f http.HandlerFunc) http.HandlerFunc {
	if !globalHTTPTrace {
		return f
	}
	return httpTracelog(f, false)
}

func httpTracelog(f http.HandlerFunc, logBody bool) http.HandlerFunc {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	name = strings.TrimPrefix(name, "github.com/minio/minio/cmd.")
	bodyPlaceHolder := []byte("<BODY>")

	return func(w http.ResponseWriter, r *http.Request) {
		const timeFormat = "2006-01-02 15:04:05 -0700"
		var reqBodyRecorder *recordRequest

		// Generate short random request ID
		reqID := fmt.Sprintf("%f", float64(time.Now().UnixNano())/1e10)

		reqBodyRecorder = &recordRequest{Reader: r.Body, logBody: logBody}
		r.Body = ioutil.NopCloser(reqBodyRecorder)

		// Setup a http response body recorder
		respBodyRecorder := &recordResponseWriter{ResponseWriter: w, logBody: logBody}

		b := bytes.NewBuffer(nil)
		fmt.Fprintf(b, "[REQUEST %s] [%s] [%s]\n", name, reqID, time.Now().Format(timeFormat))

		f(respBodyRecorder, r)

		// Build request log and write it to log file
		fmt.Fprintf(b, "%s %s", r.Method, r.URL.Path)
		if r.URL.RawQuery != "" {
			fmt.Fprintf(b, "?%s", r.URL.RawQuery)
		}
		fmt.Fprintf(b, "\n")

		fmt.Fprintf(b, "Host: %s\n", r.Host)
		for k, v := range r.Header {
			fmt.Fprintf(b, "%s: %s\n", k, v[0])
		}
		fmt.Fprintf(b, "\n")
		if logBody {
			bodyContents := reqBodyRecorder.Data()
			if bodyContents != nil {
				b.Write(bodyContents)
				fmt.Fprintf(b, "\n")
			}
		} else {
			b.Write(bodyPlaceHolder)
			fmt.Fprintf(b, "\n")
		}

		fmt.Fprintf(b, "\n")

		// Build response log and write it to log file
		fmt.Fprintf(b, "[RESPONSE] [%s] [%s]\n", reqID, time.Now().Format(timeFormat))

		b.Write(respBodyRecorder.Headers())
		fmt.Fprintf(b, "\n")
		bodyContents := respBodyRecorder.Body()
		if bodyContents != nil {
			b.Write(bodyContents)
			fmt.Fprintf(b, "\n")
		} else {
			// If there was no error response and body logging is disabled
			// then we print <BODY>
			if !logBody {
				b.Write(bodyPlaceHolder)
				fmt.Fprintf(b, "\n")
			}
		}

		fmt.Fprintf(b, "\n")

		os.Stdout.Write(b.Bytes())
	}
}
