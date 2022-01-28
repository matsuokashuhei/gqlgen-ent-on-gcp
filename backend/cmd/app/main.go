package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/resolver"
	"github.com/matsuokashuhei/landin/pkg/database"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{DB: db}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	db, _ := database.Connect()
	r := gin.Default()
	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run()
}
