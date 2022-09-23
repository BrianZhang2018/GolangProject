package mysql

import (
	"errors"
	"fmt"
	"github.com/BrianZhang2018/onsite/pkg/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type (
	SqlDb interface {
		Close() error
		GetPolls(title string) ([]model.Poll, error)
		CreatePoll(request model.PollRequest) (bool, error)
	}
	Config struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DbName   string `yaml:"dbName"`
		UserName string `yaml:"userName"`
		Pwd      string `yaml:"pwd"`
	}
	DB struct {
		sqlDb *sqlx.DB
	}
)

func (db *DB) Close() error {
	return db.sqlDb.Close()
}
func NewDBClient(cfg Config) (SqlDb, error) {
	sqlConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.UserName, cfg.Pwd, cfg.Host, cfg.Port, cfg.DbName)
	db, err := sqlx.Open("mysql", sqlConfig)
	if err != nil {
		return nil, errors.New("db connection failed")
	}
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	return &DB{
		sqlDb: db,
	}, nil
}

func (db *DB) CreatePoll(request model.PollRequest) (bool, error) {
	//tx := db.sqlDb.MustBegin()
	_, err := db.sqlDb.Exec("insert into poll (title, description, options) values (?,?,?)", request.Title, request.Description, request.Options)
	if err != nil {
		fmt.Print(err.Error())
		return false, errors.New(err.Error())
	} else {
		log.Println("created a poll, row affected %s")
	}
	//tx.Commit()
	return true, nil
}

func (db *DB) GetPolls(title string) ([]model.Poll, error) {
	//var response []model.Response
	polls := []model.Poll{}
	err := db.sqlDb.Select(&polls, "select p.* from poll p join user u on p.title=u.name")
	if err != nil {
		fmt.Print(err.Error())
		return nil, errors.New(err.Error())
	} else {
		log.Println("created a poll, row affected %s")
	}
	return polls, nil
}
