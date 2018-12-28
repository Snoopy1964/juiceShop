package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler
}

func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// log.Println("Processing request ...", r.URL)
	// log.Println("Processing request ...", r.Header)
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	encodings := r.Header.Get("Accept-Encoding")
	if !strings.Contains(encodings, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}
	w.Header().Add("Content-Encoding", "gzip")
	//w.Header().Add("Content-Type", "text/html")
	gzipwriter := gzip.NewWriter(w)
	defer gzipwriter.Close()
	var gzrw http.ResponseWriter
	if pusher, ok := w.(http.Pusher); ok {
		gzrw = gzipPushResponseWriter{
			gzipResponseWriter: gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gzipwriter,
			}, Pusher: pusher,
		}
	} else {
		gzrw = gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzipwriter,
		}
	}
	gm.Next.ServeHTTP(gzrw, r)
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

type gzipPushResponseWriter struct {
	gzipResponseWriter
	http.Pusher
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}
