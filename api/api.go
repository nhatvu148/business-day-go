package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/nhatvu148/business-day-go/db"
	"github.com/nhatvu148/business-day-go/middlewares"
	"github.com/nhatvu148/business-day-go/models"
	"github.com/nhatvu148/business-day-go/tools"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Config tools.Config
	Conn   *sql.DB
	DB     models.DBModel
}

func (app *Application) Routes() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", app.HomePageHandler)
	r.HandleFunc("/catfact", app.CatFactPageHandler)
	r.HandleFunc("/api/v1/business-day", app.BusinessDayHandler)
	r.HandleFunc("/api/v1/custom-holiday", app.CustomHolidayHandler)

	m := middlewares.RequestPathLogger(middlewares.SetCors(r))

	return m
}

func (app *Application) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.Config.Port),
		Handler:           app.Routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Info().Msg(fmt.Sprintf("Starting HTTP server in %s mode on port %s\n", app.Config.Env, app.Config.Port))
	return srv.ListenAndServe()
}

func SetupApp() *Application {
	config := tools.Config{}
	config.LoadConfig()
	config.SetLogger()

	conn := db.InitDB(config.DatabaseURL)

	db.RunDBMigration(config.MigrationURL, config.DatabaseURL)

	app := &Application{
		Config: config,
		Conn:   conn,
		DB:     models.DBModel{DB: conn},
	}
	return app
}

func Run() {
	app := SetupApp()
	defer app.Conn.Close()

	err := app.Serve()
	if err != nil {
		log.Fatal().Err(err).Msg("Server connection error")
	}
}
