package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/github-real-lb/bookings-web-app/internal/models"
	"github.com/github-real-lb/bookings-web-app/util"
)

// app holds the configurations and templates of the app
var app util.AppConfig

func main() {
	err := InitializeApp(util.DevelopmentMode)
	if err != nil {
		log.Fatal(err)
	}
	server := NewServer(app.ServerAddress)

	// start the server
	fmt.Println("Starting web server on,", app.ServerAddress)
	err = server.ListenAndServe()
	log.Fatal(err)
}

// InitializeApp loads the app configurations and setup based on the application mode
func InitializeApp(appMode util.AppMode) error {
	var err error

	// load application configurations and set app to developement mode
	app = util.LoadConfig()

	switch appMode {
	case util.DevelopmentMode:
		app.SetDevelopementMode()
	case util.TestingMode:
		app.SetTestingMode()
	default:
	}

	// load templates cache to AppConfig
	app.TemplateCache, err = GetTemplatesCache()
	if err != nil {
		return errors.New(fmt.Sprint("error creating gohtml templates cache: ", err.Error()))
	}

	// setting up session manager
	session := scs.New()
	session.Lifetime = 24 * time.Hour // keeps session data for 24 hours
	session.Cookie.Persist = true     // keeps session data after browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProductionMode() // determines use of SSL encryption
	app.Session = session

	// defining session stored types
	gob.Register(models.Reservation{})

	return nil
}
