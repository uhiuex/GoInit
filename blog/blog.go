package blog

import (
	"fmt"
	"net/http"

)

func Blog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Hey Blog!")
}