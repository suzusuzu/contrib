package sessions

import (
	"github.com/dsoprea/goappenginesessioncascade"
	"github.com/gorilla/sessions"
)

type AppEngineStore interface {
	Store
}

func NewAppEngineStore(backEnd int, keyPairs ...[]byte) AppEngineStore {
	appenginestore := cascadestore.NewCascadeStore(backEnd, keyPairs...)
	return &appEngineStore{appenginestore}
}

type appEngineStore struct {
	*cascadestore.CascadeStore
}

func (c *appEngineStore) Options(options Options) {
	c.CascadeStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
