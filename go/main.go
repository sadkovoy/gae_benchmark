package main

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

type Human struct {
	FirstName string
	LastName  string
	Company   string
}

func main() {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, "wixgamma")
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/_ah/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/benchmark", func(c *gin.Context) {
		keys := []*datastore.Key{
			datastore.NameKey("Human", "ID1", nil),
			datastore.NameKey("Human", "ID2", nil),
			datastore.NameKey("Human", "ID3", nil),
		}
		humans := make([]Human, 3)

		if err := dsClient.GetMulti(ctx, keys, humans); err != nil {
			c.JSON(500, gin.H{"error": err})
		}

		c.JSON(200, gin.H{"humans": humans})
	})

	r.Run()
}
