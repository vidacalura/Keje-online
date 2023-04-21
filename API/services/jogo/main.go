package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sort"
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
	Tabuleiro  [5][5]string `json:"tabuleiro"`
	Turno      string       `json:"turno"`
	Movimentos []Movimento  `json:"movimentos"`
}

type Jogador struct {
	SocketId string `json:"socketId"`
	Username string `json:"username"`
	Lado     string `json:"lado"`
	Pontos   int    `json:"pontos"`
	Tempo    int    `json:"tempo"`
}

//type espectador struct {}

type Movimento struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Lado string `json:"lado"`
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
		jogo.POST("/salas", criarSala)
		jogo.POST("/salas/entrar", entrarSala)
		jogo.POST("/", fazerMovimento)
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
	for i, sala := range salas {
		if sala.Id == 0 {
			salas[i].Id = gerarCodigoSala()
			salas[i].Jogo.Turno = "B"

			c.IndentedJSON(http.StatusOK, gin.H{ "salaId": salas[i].Id, "message": "Sala criada com sucesso!" })
			return
		}
	}  

	c.IndentedJSON(http.StatusOK, gin.H{ "error": "Não foi possível criar uma sala. Tente novamente mais tarde." })
}

func entrarSala(c *gin.Context) {
	type ReqBody struct {
		SalaId   int    `json:"salaId"`
		SocketId string `json:"socketId"`
		Username string `json:"username"`
		Tempo    int    `json:"tempo"`
	}

	var reqBody ReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Dados inválidos." })
		return
	}

	if reqBody.SocketId == "" || reqBody.Tempo == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Dados insuficientes." })
		return
	}

	for i, sala := range salas {
		if sala.Id == reqBody.SalaId {
			var jogador Jogador
			jogador.SocketId = reqBody.SocketId
			jogador.Username = reqBody.Username
			jogador.Tempo = reqBody.Tempo

			if sala.Jogadores[0].SocketId == "" {
				jogador.Lado = sortearLado()
				salas[i].Jogadores[0] = jogador
			} else if sala.Jogadores[1].SocketId == "" {
				if sala.Jogadores[0].Lado == "B" {
					jogador.Lado = "P"
				} else {
					jogador.Lado = "B"
				}

				salas[i].Jogadores[1] = jogador
			} else {
				c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "A sala está cheia." })
				return
			}

			c.IndentedJSON(http.StatusOK, gin.H{ "message": "Você entrou na sala com sucesso!" })
			return
		}
	}  

	c.IndentedJSON(http.StatusNotFound, gin.H{ "error": "Sala não encontrada." })
}

func fazerMovimento(c *gin.Context) {
	type ReqBody struct {
		SalaId     int         `json:"salaId"`
		SocketId   string      `json:"socketId"`
		Movimentos []Movimento `json:"movimentos"`
	}

	var reqBody ReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Dados inválidos." })
		return
	}

	for i, sala := range salas {
		if reqBody.SalaId == sala.Id {
			// Validação de usuário
			// Checar se é vez do usuário

			// Validação de movimento
			if isConnected(reqBody.Movimentos, sala.Jogo.Tabuleiro) {
				for _, movimento := range reqBody.Movimentos {
					salas[i].Jogo.Tabuleiro[movimento.Y][movimento.X] = movimento.Lado
				}

				// Validar se jogo acabou
				

				// Passar vez
				if salas[i].Jogo.Turno == "B" {
					salas[i].Jogo.Turno = "P"
				} else {
					salas[i].Jogo.Turno = "B"
				}

				c.IndentedJSON(http.StatusOK, gin.H{ "sala": salas[i], "message": "Movimento realizado com sucesso!" })
				return
			} else {
				c.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Movimento inválido." })
				return
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "error": "Sala não encontrada." })
}

func isConnected(movimento []Movimento, tabuleiro [5][5]string) bool {
	if movimento == nil {
		return false
	}

	if len(movimento) > 3 || len(movimento) < 1 { // Possível mudança para 4 peças por rodada
		return false
	}

	for _, m := range movimento {
		if m.X > 4 || m.X < 0 || m.Y > 4 || m.Y < 0 {
			return false
		}
	}

	// Verifica lance espelhado

	// Verifica se casa já foi preenchida
	for _, m := range movimento {
		if tabuleiro[m.Y][m.X] != "" {
			return false
		}
	}

	if len(movimento) == 1 {
		return true
	}

	// Verifica se o lance é legal
	var eixoComum Movimento 
	for _, m := range movimento {
		if eixoComum.Lado == "" {
			eixoComum = m
		} else {
			if m.X == eixoComum.X {
				eixoComum.Y = -1
			} else if m.Y == eixoComum.Y {
				eixoComum.X = -1
			} else {
				return false
			}
		}
	}

	// Verifica se peças estão conectadas
	var sortedMovimentos []int
	if eixoComum.X == -1 {
		for _, m := range movimento {
			sortedMovimentos = append(sortedMovimentos, m.X)
		}
	} else {
		for _, m := range movimento {
			sortedMovimentos = append(sortedMovimentos, m.Y)
		}
	}

	sort.Ints(sortedMovimentos)
	log.Println(sortedMovimentos)
	for i := 1; i < len(sortedMovimentos); i++ {
		if sortedMovimentos[i] != sortedMovimentos[i - 1] + 1 {
			return false
		}
	}

	return true
}

func countCasas(tabuleiro [5][5]string) int {
	countCasas := 0
	for _, linha := range tabuleiro {
		for _, casa := range linha {
			if casa != "" {
				countCasas++
			}
		}
	}

	return countCasas
}

func restartJogo(sala *Sala) {
	if sala.Jogadores[0].Lado == "B" {
		sala.Jogadores[0].Lado == "P"
		sala.Jogadores[1].Lado == "B"
	} else {
		sala.Jogadores[0].Lado == "B"
		sala.Jogadores[1].Lado == "P"
	}

	sala.Jogadores[0].Tempo = 180
	sala.Jogadores[1].Tempo = 180

	sala.Jogo.Tabuleiro = [5][5]string{}
	sala.Jogo.Turno = "B"
	sala.Jogo.Movimentos = []Movimento{}

	sala.JogoEncerrado = false
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

func sortearLado() string {
	randNum := rand.Intn(2)

	if randNum == 0 {
		return "B"
	}

	return "P"
}