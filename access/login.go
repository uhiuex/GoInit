package login

import (
	"fmt"
	"net/http"

)

func Loginhandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Hey Login Works!")
}