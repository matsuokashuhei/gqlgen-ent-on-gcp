package main

import (
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/resolver"
	"github.com/matsuokashuhei/landin/internal/middleware"
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
	r.Use(middleware.GinContextToContext())
	r.Use(middleware.Auth(db))
	r.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	r.POST("/query", graphqlHandler(db))
	r.GET("/", playgroundHandler())
	r.Run()
}
