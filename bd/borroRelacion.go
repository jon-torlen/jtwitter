package bd

import (
	"time"
	"context"
	"github.com/jon-torlen/jtwitter/models"
)

/*BorroRelacion borra la relacion en la BD*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("jtwitter")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}