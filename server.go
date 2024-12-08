package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mrandiw/go-graphql-simple/config"
	"github.com/mrandiw/go-graphql-simple/graph"
	"github.com/mrandiw/go-graphql-simple/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	// Initialize Gin
	r := gin.Default()

	r.SetTrustedProxies(nil)

	// Initialize PostgreSQL connection
	db, err := config.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// GraphQL Handler
	r.POST("/query", func(c *gin.Context) {
		graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))

		// Handle the GraphQL request
		graphqlHandler.ServeHTTP(c.Writer, c.Request)

		// Add custom response if needed
		if c.Writer.Status() >= http.StatusBadRequest {
			utils.CustomResponse(c, c.Writer.Status(), "Error occurred while processing GraphQL request", nil)
		}
	})

	// Playground Handler
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	log.Println("Server is running on http://localhost:" + port)
	r.Run("127.0.0.1:" + port)
}
