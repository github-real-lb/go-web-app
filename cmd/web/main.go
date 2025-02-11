package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/github-real-lb/bookings-web-app/db"
	"github.com/github-real-lb/bookings-web-app/util/config"
	"github.com/github-real-lb/bookings-web-app/util/loggers"
	"github.com/github-real-lb/bookings-web-app/util/mailers"
)

func main() {
	// initialize application
	err := InitializeApp(config.DevelopmentMode)
	if err != nil {
		log.Fatal("Error initializing application:", err)
	}

	// create a new database connection pool
	dbStore, err := db.NewPostgresDBStore(app.DBConnectionString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer dbStore.(*db.PostgresDBStore).DBConnPool.Close()

	// create a new error logger
	errLogger := loggers.NewSmartLogger(nil, "ERROR\t")

	// create a new info logger
	infoLogger := loggers.NewSmartLogger(nil, "INFO\t")

	// create a new mailer
	mailer := mailers.NewSmartMailer()

	// create a new server
	server := NewServer(dbStore, errLogger, infoLogger, mailer)

	// load web page templates cache
	err = server.Renderer.LoadGoHtmlPageTemplates()
	if err != nil {
		log.Fatal(fmt.Sprint("error creating gohtml web page templates cache: ", err.Error()))
	}

	// load mail templates cache
	err = server.Renderer.LoadGoHtmlMailTemplates()
	if err != nil {
		log.Fatal(fmt.Sprint("error creating gohtml mail templates cache: ", err.Error()))
	}

	// start server in a separate goroutine
	go server.Start()

	// Listen for interrupt signal (Ctrl+C) or SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)    // Ctrl+C
	signal.Notify(stop, syscall.SIGTERM) // SIGTERM
	defer close(stop)

	// block until a stop signal is received
	<-stop

	// stop server
	server.Stop()
}
