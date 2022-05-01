package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/resolver"
	"github.com/matsuokashuhei/landin/internal/middleware"
)

const defaultPort = "8080"

func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	h := handler.NewDefaultServer(resolver.NewSchema(client))
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
	// db, _ := database.Connect()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	log.Print(dsn)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("opening ent client: %v", err)
	}
	defer client.Close()

	r := gin.Default()
	r.Use(middleware.GinContextToContext())
	r.Use(middleware.Auth(client))
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
	r.POST("/query", graphqlHandler(client))
	r.GET("/", playgroundHandler())
	r.Run()
}
