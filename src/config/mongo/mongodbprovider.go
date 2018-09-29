package mongo

import (
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

type SessionProvider interface {
	GetSession() *mgo.Session
}
type sessionProvider struct {
	session *mgo.Session
}

var MongoSessionProvider SessionProvider

func (sessionProvider sessionProvider) GetSession() *mgo.Session {
	return sessionProvider.session.Clone()
}
func init() {
	host := viper.GetString("databaseHost") //localhost
	sessionAux, err := mgo.Dial(host)
	if err == nil {
		aux := &sessionProvider{sessionAux}

		MongoSessionProvider = aux
	} else {
		//TODO: LOG ERROR
	}

}
