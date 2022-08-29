// Copyright (c) 2020 Changkun Ou. All rights reserved.

package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

//go:embed public/*
var public embed.FS

const (
	addr = "0.0.0.0:3001"
)

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	logger := logging(l)

	fsys, err := fs.Sub(public, "public")
	if err != nil {
		l.Fatalf("cannot access sub file system: %v", err)
	}

	r := http.NewServeMux()
	r.Handle("/", http.FileServer(http.FS(fsys)))

	r.HandleFunc("/root.txt", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("b95a7a7638225f85ebfbfbcc69adc284"))
		resp.Header().Set("Content-Type", "text/plain")
	})

	s := &http.Server{
		Addr:         addr,
		Handler:      logger(r),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		l.Println("marlon.life is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.SetKeepAlivesEnabled(false)
		if err := s.Shutdown(ctx); err != nil {
			l.Fatalf("cannot gracefully shutdown marlon.life: %v", err)
		}
		close(done)
	}()

	l.Printf("marlon.life is serving on http://%s", addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatalf("cannot listen on %s, err: %v\n", addr, err)
	}

	l.Println("goodbye!")
	<-done
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				logger.Println(readIP(r), r.Method, r.URL.Path)
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func readIP(r *http.Request) string {
	clientIP := r.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}
	if addr := r.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
		return addr
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return "unknown" // use unknown to guarantee non empty string
	}
	return ip
}
