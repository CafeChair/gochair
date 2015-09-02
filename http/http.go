package http
import (
	"x/comm"
	"net/http"
)

func InfoRouter() {
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte(comm.VERSION))
	})
}
