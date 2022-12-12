package middlew

import (
	"github.com/jon-torlen/jtwitter/bd"
	"net/http"
)

/*ChequeoBD es el middlew que me permite saber el estado de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
