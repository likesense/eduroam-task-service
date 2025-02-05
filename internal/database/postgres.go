package postgres

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresCred struct {
	host     string
	port     int
	dbName   string
	username string
	password string
}

var connection *postgresCred

func getPostgresCredInstance() *postgresCred {
	if connection == nil {
		port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		if err != nil {
			fmt.Println("Environment variable 'POSTGRES_PORT' must be type of INT and not nil")
		}
		connection = &postgresCred{
			host:     os.Getenv("TASK_POSTGRES_HOST"),
			port:     port,
			dbName:   os.Getenv("TASK_POSTGRES_DB_NAME"),
			username: os.Getenv("TASK_POSTGRES_USERNAME"),
			password: os.Getenv("TASK_POSTGRES_PASSWORD"),
		}
	}
	return connection
}

func (pgc *postgresCred) String() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pgc.host, pgc.port, pgc.username, pgc.password, pgc.dbName,
	)
}

func NewPostgresDbConnection() *sqlx.DB {
	db, err := sqlx.Connect(os.Getenv("POSTGRES_DRIVER"), getPostgresCredInstance().String())
	if err != nil {
		log.Fatalf("Can't connect to Postgres DB (Redo): %s\n", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error during Postgres DB (Redo) connection check: %s\n", err.Error())
	}

	log.Println("Connection to Postgres DB (Redo) successfully established")

	return db
}
