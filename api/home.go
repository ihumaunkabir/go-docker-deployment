package home

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"Greetings" : "You are awesome!"}`))
}
