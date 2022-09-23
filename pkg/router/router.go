package router

import (
	"github.com/BrianZhang2018/onsite/internal/config"
	"github.com/BrianZhang2018/onsite/pkg/api"
	"github.com/BrianZhang2018/onsite/pkg/object/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

type Config struct {
	MySqlClient mysql.SqlDb
	Config      config.AppConfig
}

func New(cfg Config) *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/health").Handler(api.AppError(api.Health))

	testService := api.TestService{
		Config: cfg.Config,
		Db:     cfg.MySqlClient,
	}
	r.Methods(http.MethodPost).Path("/createPoll").Handler(api.AppError(testService.CreatePoll))
	r.Methods(http.MethodPost).Path("/getPolls").Handler(api.AppError(testService.GetPolls))

	return r
}
