package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Human struct {
	FirstName string
	LastName  string
	Company   string
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/_ah/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/benchmark", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)

		keys := []*datastore.Key{
			datastore.NewKey(ctx, "Human", "ID1", 0, nil),
			datastore.NewKey(ctx, "Human", "ID2", 0, nil),
			datastore.NewKey(ctx, "Human", "ID3", 0, nil),
		}
		humans := make([]Human, 3)

		if err := datastore.GetMulti(ctx, keys, humans); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"humans": humans})
		}
	})

	http.Handle("/", r)
	appengine.Main()
}
