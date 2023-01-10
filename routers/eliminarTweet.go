package routers

import (
	"github.com/jon-torlen/jtwitter/bd"
	"net/http"
)

/*EliminarTweet permite borrar un Tweet determinado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1 {
		http.Error(w, "Debe enviar el parametro ID",http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el Tweet"+err.Error(	),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}