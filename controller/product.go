package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Product struct{}

func (pr Product) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	WriteJSON(w, http.StatusOK, "OPEN FOR PUBLIC")
}

func (pr Product) Auth(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	WriteJSON(w, http.StatusOK, "NEED LOGIN")
}

func (pr Product) SuperAdmin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	WriteJSON(w, http.StatusOK, "OPEN ONLY FOR SUPER ADMIN")
}
