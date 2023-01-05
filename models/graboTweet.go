package models

import (
	"time"
)
/*GraboTweet es el formato de estructura que tendra nuestro Tweet*/
type GraboTweet struct {
	UserID string `bson:"userid" json:"userid,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
}