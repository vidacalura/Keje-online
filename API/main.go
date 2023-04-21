// API Gateway do sistema do jogo
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	jogo := r.Group("/api/jogo")
	{
		jogo.GET("/:idSala", jogoHandler)
		jogo.GET("/analise/:codJogo", jogoHandler)
		jogo.POST("/", jogoHandler)
	}

	users := r.Group("/api/usuarios")
	{
		users.GET("/ping", usersHandler)
	}

	auth := r.Group("/api/auth")
	{
		auth.GET("/ping", authHandler)
	}

	r.Run(os.Getenv("APIGateway"))
}

func jogoHandler(c *gin.Context) {
	reqUrl, err := url.Parse(os.Getenv("JogoMS"))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ "error": "Erro ao conectar com o servidor. Tente novamente mais tarde." })
		return
	}

	redirecionarRequest(reqUrl, c)
}

func usersHandler(c *gin.Context) {
	reqUrl, err := url.Parse(os.Getenv("UserMS"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ "error": "Erro ao conectar com o servidor. Tente novamente mais tarde." })
		return
	}

	redirecionarRequest(reqUrl, c)
}

func authHandler(c *gin.Context) {
	reqUrl, err := url.Parse(os.Getenv("AuthMS"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ "error": "Erro ao conectar com o servidor. Tente novamente mais tarde." })
		return
	}

	redirecionarRequest(reqUrl, c)
}

func redirecionarRequest(reqUrl *url.URL, c *gin.Context) {
	proxy := httputil.NewSingleHostReverseProxy(reqUrl)
	proxy.ServeHTTP(c.Writer, c.Request)
}