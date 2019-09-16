package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		logrus.WithField("req", r).Info("income request")
		w.Write([]byte("world"))
	})
	http.ListenAndServe(":8080", nil)
}
