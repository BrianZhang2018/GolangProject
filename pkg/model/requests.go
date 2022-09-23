package model

import "database/sql"

type Request struct {
	Id int64 `json:"id,omitempty"`
}

type Response struct {
	Id   int64  `db:"id,omitempty"`
	Name string `db:"name,omitempty"`
}

type Poll struct {
	PollId      int64           `db:"pollId,omitempty"`
	Title       string          `db:"title,omitempty"`
	Description string          `db:"description,omitempty"`
	Options     string          `db:"options,omitempty"`
	UrlLink     *sql.NullString `db:"urlLink,omitempty"`
	Ts          string          `db:"ts,omitempty"`
}

type PollResponse struct {
	Id   int64  `db:"id,omitempty"`
	Name string `db:"name,omitempty"`
}

type User struct {
	Id   int64  `db:"id,omitempty"`
	Name string `db:"name,omitempty"`
}

type PollRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Options     string `json:"options,omitempty"`
}
