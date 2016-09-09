package mainappengine

import (
	"github.com/dsoprea/goappenginesessioncascade"
	"github.com/gin-gonic/gin"
	"github.com/suzusuzu/contrib/sessions"
	"net/http"
)

func init() {
	r := gin.Default()
	store := sessions.NewAppEngineStore(cascadestore.AllBackends, []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count += 1
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	http.Handle("/", r)
}
