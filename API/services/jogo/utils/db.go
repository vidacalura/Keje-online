// db.go contém todas as funções relacionadas à conexão
// com o banco de dados
package utils

import (
	//"database/sql"
	"log"
	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConectarBD() { //*sql.DB {
	// Pega dados de .env
	err := godotenv.Load("ENV")
	if err != nil {
		log.Fatal(err)
	}

	// Conecta ao banco de dados
	//db, err := sql.Open("mysql", os.Getenv("DSN"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Checa conexão com o banco de dados
	//if err = db.Ping(); err != nil {
	//	log.Fatal(err)
	//}

	//return db
} 