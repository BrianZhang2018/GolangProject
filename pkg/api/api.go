package api

import (
	"encoding/json"
	"github.com/BrianZhang2018/onsite/internal/config"
	"github.com/BrianZhang2018/onsite/pkg/model"
	"github.com/BrianZhang2018/onsite/pkg/object/mysql"
	"net/http"
)

type TestService struct {
	Config config.AppConfig
	Db     mysql.SqlDb
}

type TestRequest struct {
	Id string
}

type TestResponse struct {
	Res bool
}

type PollResponse struct {
	Polls []model.Poll
}

func (test *TestService) CreatePoll(w http.ResponseWriter, r *http.Request) error {
	var request model.PollRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return &MyError{DeveloperText: "invalid request"}
	}
	//TODO: request validation

	res, _ := test.Db.CreatePoll(request)
	writeResponse(w, http.StatusOK, &TestResponse{res})
	return nil
}

func (test *TestService) GetPolls(w http.ResponseWriter, r *http.Request) error {
	title := r.URL.Query().Get("title")
	if title == "" {
		return &MyError{DeveloperText: "missing title in request"}
	}
	//TODO: request validation
	res, _ := test.Db.GetPolls(title)
	writeResponse(w, http.StatusOK, &PollResponse{res})
	return nil
}
