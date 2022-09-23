package main

import (
	"github.com/BrianZhang2018/onsite/internal/config"
	"github.com/BrianZhang2018/onsite/pkg/object/mysql"
	"github.com/BrianZhang2018/onsite/pkg/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	config, err := config.GetAppConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbClient, err := mysql.NewDBClient(config.DbConfig)
	if err != nil {
		log.Fatal("db connection failed")
	}

	defer dbClient.Close()

	routerConfig := router.Config{
		Config:      *config,
		MySqlClient: dbClient,
	}
	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "X-Auth-Token", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Info("Application Started Up")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router.New(routerConfig))))
}
