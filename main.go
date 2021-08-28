package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/uhiuex/hello/access"

)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Hello World!")
}

func blog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Hey Blog!")
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/blog", blog)
	r.HandleFunc("/access", login)
	fmt.Println("Loading Server....")
	http.ListenAndServe(":8000", r)
}


