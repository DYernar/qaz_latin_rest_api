package controller

import "net/http"

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("appToken") != AppToken {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
