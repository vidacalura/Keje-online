package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"jeek2-ms-jogo/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Sala struct {
	Id            int        `json:"id"`
	Jogadores     [2]Jogador `json:"jogadores"`
	Jogo          Jogo       `json:"jogo"`
	// Espectadores []espectador `json:"espectadores"`
	JogoEncerrado bool       `json:"jogoEncerrado"`
}

type Jogo struct {
	Tabuleiro  [5][5]byte  `json:"tabuleiro"`
	Turno      byte        `json:"turno"`
	Movimentos []Movimento `json:"movimentos"`
}

type Jogador struct {
	SocketId string `json:"socketId"`
	Username string `json:"username"`
	Lado     byte   `json:"lado"`
	Pontos   int    `json:"pontos"`
	Tempo    int    `json:"tempo"`
}

//type espectador struct {}

type Movimento struct {
	X    int  `json:"x"`
	Y    int  `json:"y"`
	Lado byte `json:"lado"`
}

var db *sql.DB
var salas [100]Sala

func main() {
	//db = utils.ConectarBD()
	utils.ConectarBD()

	r := gin.Default()

	jogo := r.Group("/api/jogo")
	{
		jogo.GET("/:idSala", getSala)          // Jogo presente (API)
		jogo.GET("/analise/:codJogo", getJogo) // Jogo passado (BD)
		jogo.POST("/", criarSala)
		//jogo.POST("/", fazerMovimento)
	}

	r.Run(os.Getenv("JogoMS"))
}

func getSala(c *gin.Context) {
	idSalaStr := c.Param("idSala")
	idSala, err := strconv.Atoi(idSalaStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Id de sala inválido." })
		return
	}

	for _, sala := range salas {
		if sala.Id == idSala {
			c.IndentedJSON(http.StatusOK, gin.H{ "sala": sala, "message": "Sala encontrada com sucesso!" })
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "error": "Sala não encontrada." })
}

func getJogo(c *gin.Context) {
	//codJogo := c.Param("codJogo")

}

func criarSala(c *gin.Context) {
	var jogador1 Jogador
	// recebe: SocketId, Username, Tempo
	if err := c.BindJSON(&jogador1); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Dados inválidos." })
		return
	}

	if jogador1.SocketId == "" || jogador1.Tempo == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Dados insuficientes." })
		return
	}

	jogador1.Lado = sortearLado()

	for i, sala := range salas {
		if sala.Id == 0 {
			salas[i].Id = gerarCodigoSala()
			salas[i].Jogadores[0] = jogador1
			salas[i].Jogo.Turno = 'B'

			c.IndentedJSON(http.StatusOK, gin.H{ "sala": salas[i], "message": "Sala criada com sucesso!" })
			return
		}
	}  

	c.IndentedJSON(http.StatusOK, gin.H{ "error": "Não foi possível criar uma sala. Tente novamente mais tarde." })
}

func gerarCodigoSala() int {
	codSala := (rand.Intn(9) + 1) * 100000 +
				rand.Intn(10) * 10000 +
				rand.Intn(10) * 1000 +
				rand.Intn(10) * 100 +
				rand.Intn(10) * 10 +
				rand.Intn(10)

	// Validar se código já existe
	for _, sala := range salas {
		if sala.Id == codSala {
			return gerarCodigoSala()
		}
	}

	return codSala
}

func sortearLado() byte {
	randNum := rand.Intn(2)

	if randNum == 0 {
		return 'B'
	}

	return 'P'
}