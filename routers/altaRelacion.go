package routers

import (
	"github.com/jon-torlen/jtwitter/bd"
	"github.com/jon-torlen/jtwitter/models"
	"net/http"
)

/*AltaRelacion realiza el registro de la relacionentre usuarios */
func AltaRelacion(w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")
	if len(ID)<1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
if err != nil {
	http.Error(w, "El parametro ID es obligatorio"+err.Error(),http.StatusBadRequest)
	return
}

if status == false {
	http.Error(w, "No se ha logradoinsertar la  relacion "+err.Error(),http.StatusBadRequest)
	return
}
w.WriteHeader(http.StatusCreated)

}