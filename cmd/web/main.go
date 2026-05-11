package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

const port string = ":8181"

func main() {

	addr := flag.String("addr", port, "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// This logger will show the source file and line of code of where the log
	// was generated. I see usefulness for debugging.
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// }))

	db, err := sql.Open("mysql", "web:pass@snippetbox?parseTime=true")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
		// Could also just do this:
		// logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	logger.Info("Server starting on localhost", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
