package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/sniperkit/yaag/middleware"
	"github.com/sniperkit/yaag/yaag"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().String())
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = ioutil.ReadAll(r.Body)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("test", "tesasasdasd")
	fmt.Fprintf(w, time.Now().String())
}

func main() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Gorilla Mux", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/testing", postHandler).Methods("POST")
	http.ListenAndServe(":8080", middleware.Handle(r))
}
