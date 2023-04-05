package controllers

import "net/http"

func (*Controllers) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is comming from a private end-point!"))
	}
}
