package handlers

import "net/http"

func authorized(next http.HandlerFunc) http.HandlerFunc {

	return http.HandleFunc(func(w http.ResponseWriter,r *http.Request){
		
	})
}
