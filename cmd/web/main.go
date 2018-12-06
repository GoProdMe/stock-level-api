package main

import (
	"database/sql"
	"flag"
	"log"
	"stock-level-api/cmd/models"

	_ "github.com/go-sql-driver/mysql" // init() function required, so alias with blank identifier
)

func main() {

	// Define command-line flags for the network address and location of the static
	// files directory.
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	dsn := flag.String("dsn", "root:@/stock_api?parseTime=true", "MySQL DSN") // no password
	tlsCert := flag.String("tls-cert", "./tls/cert.pem", "Path to TLS certificate")
	tlsKey := flag.String("tls-key", "./tls/key.pem", "Path to TLS key")

	// parse the command-line flags.
	flag.Parse()

	db := connect(*dsn)

	defer db.Close()

	// Initialize a new instance of App containing the dependencies.
	app := &App{
		Addr:     *addr,
		Database: &models.Database{db},
		TLSCert:  *tlsCert,
		TLSKey:   *tlsKey,
	}

	app.RunServer()
}

// The connect() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Set the maximum number of idle connections in the pool. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
