package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/tariq1890/stock-ticker/pkg/config"
	"github.com/tariq1890/stock-ticker/pkg/stocks"
)

type stocksHandler struct {
	service stocks.Service
	ndays   int
	symbol  string
}

func (s *stocksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := s.service.GetDataFor(s.symbol, s.ndays)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("Unexpected error occurred: %v", err)
	}
	jsonResp, _ := json.MarshalIndent(p, "", "   ")
	w.Write(jsonResp)
	return
}

func main() {
	cfg := config.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	stockSvc := stocks.NewService(cfg.APIURL, cfg.APIKey)

	r := mux.NewRouter()
	r.Handle("/api/stockticker", &stocksHandler{stockSvc, int(cfg.NDays), cfg.Symbol})
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
