package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/uhiuex/goinit/access"
	"github.com/uhiuex/goinit/blog"

)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Hello World!")
}

func main(){
	r := mux.NewRouter()
	pt := r.NewRoute().PathPrefix("/app").Subrouter()
	r.HandleFunc("/", home)
	pt.HandleFunc("/blog", blog.Blog)
	pt.HandleFunc("/login", login.Loginhandler).Methods("GET")
	fmt.Println("Loading Server....")
	http.ListenAndServe(":8000", r)

	http.Handle("/", r)
}


