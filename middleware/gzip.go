package middleware

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler
}

func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("GzipMiddleware - Request URL: %v", r.URL)
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	encodings := r.Header.Get("Accept-Encoding")
	if !strings.Contains(encodings, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}
	w.Header().Add("Content-Encoding", "gzip")
	gzipwriter := gzip.NewWriter(w)
	defer gzipwriter.Close()
	var gzrw http.ResponseWriter
	if pusher, ok := w.(http.Pusher); ok {
		log.Printf("GzipMiddleware - pusher called: %v", r.URL)
		gzrw = gzipPushResponseWriter{
			gzipResponseWriter: gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gzipwriter,
			},
			Pusher: pusher,
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
