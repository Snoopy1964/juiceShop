package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("TimeoutMiddleware - Request URL: %v", r.URL)
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	r.WithContext(ctx)
	ch := make(chan struct{})
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return
	case <-ctx.Done():
		log.Printf("TimeoutMiddleware - Timeout Request: %v", r.Header)
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
