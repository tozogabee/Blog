package handlers

import (
	"database/sql"
	"net/http"
)

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}

func (h handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get All Comments"))
}

func (h handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create comment"))
}
