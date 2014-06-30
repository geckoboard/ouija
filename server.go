package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type HTTPServer struct {
	Dist *Distribution
}

func (s HTTPServer) ListenAndServe(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		latency := s.Dist.OutputRand()
		log.Println(r.Method, r.URL, latency)

		// Convert latency to nanoseconds because that's what time.Sleep() needs.
		latencyNs := latency * float64(time.Second)
		time.Sleep(time.Duration(latencyNs))

		w.Header().Set("X-Ouija-Latency", fmt.Sprintf("%v", latency))
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("YO\n"))
	})

	http.Serve(l, handler)
	return nil
}
