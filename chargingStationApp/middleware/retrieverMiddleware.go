package middleware

import (
	stations "charging-assignment1/handlers"
	"context"
	"net/http"
)

var retriever stations.ChargingInterface = &stations.Database{}

func RetrieverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var retriever stations.ChargingInterface

		if r.Header.Get("Q-Testing") == "true" {
			retriever = &stations.Mock{}
		} else {
			retriever = &stations.Database{}
		}

		// Store the retriever object in the request context
		ctx := context.WithValue(r.Context(), "retriever", retriever)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
