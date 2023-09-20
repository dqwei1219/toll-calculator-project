package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/dqwei1219/toll-calculator-project/types"
	"google.golang.org/grpc"
)

func main() {
	httplistenAddr := flag.String("httpAddr", ":3000", "HTTP server listen address")
	grpclistenAddr := flag.String("grpcAddr", ":3001", "GRPC server listen address")
	flag.Parse()

	var (
		store = NewInMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLogMiddleware(svc)
	// boot up both transport together with one in a goroutine
	go makeGRPCtransport(*grpclistenAddr, svc)
	makeHTTPtransport(*httplistenAddr, svc)
}

func makeGRPCtransport(listenAddr string, svc Aggregator) error {
	// make a TCP listeners
	fmt.Printf("listening on %s\n", listenAddr)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	// close the go routine when the function returns
	defer listener.Close()
	// Make a new GRPC native server with (options)
	server := grpc.NewServer([]grpc.ServerOption{}...)

	// Register the server with the GRPC package
	types.RegisterDistAggregatorServer(server, NewGRPCServer(svc))

	return server.Serve(listener)
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
