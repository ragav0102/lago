package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
)

func Respond200(w http.ResponseWriter, r *http.Request, b []byte) {
	w.Header().Set("ReqId", strconv.FormatInt(rand.Int63n(10000000000), 10))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Respond404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ReqId", strconv.FormatInt(rand.Int63n(10000000000), 10))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}
