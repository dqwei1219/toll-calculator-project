package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dqwei1219/toll-calculator-project/types"
)

func main() {
	listenAddr := flag.String("listen-addr", ":3000", "server listen address")
	flag.Parse()

	var (
		store = NewInMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLogMiddleware(svc)
	makeHTTPtransport(*listenAddr, svc)
}

func makeHTTPtransport(listenAddr string, svc Aggregator) {
	fmt.Printf("listening on %s\n", listenAddr)
	http.HandleFunc("/aggregate", handleAggreagtor(svc))
	http.HandleFunc("/invoices", handleInvoices(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleInvoices(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get("unitId")
		if value == "" {
			writeJSON(w,
				http.StatusBadRequest,
				map[string]string{"error": "unitId is required"},
			)
		}
		unitId, err := strconv.Atoi(value)
		if err != nil {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": "unitId must be an integer"},
			)
		}
		invoice, err := svc.CalculateInvoice(unitId)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, invoice)
	}
}

func handleAggreagtor(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func writeJSON(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}
