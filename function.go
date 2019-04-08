package function

import "net/http"

func F(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	msg := []byte("Hello world")
	w.Write(msg)
}
