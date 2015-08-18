package main

import (
	"github.com/zenazn/goji/web"
//	"github.com/wcl48/valval"

	"net/http"
	"encoding/base64"
	"strings"
//	"html/template"
//	"github.com/syo-sa1982/crudstudy/models"
//	"fmt"
//	"strconv"

)

const Password = "user:user"

func UserIndex(c web.C, w http.ResponseWriter, r *http.Request) {
}

func UserNew(c web.C, w http.ResponseWriter, r *http.Request){
}

func UserCreate(c web.C, w http.ResponseWriter, r *http.Request){
}

func UserEdit(c web.C, w http.ResponseWriter, r *http.Request){
}

func UserUpdate(c web.C, w http.ResponseWriter, r *http.Request){
}

func UserDelete(c web.C, w http.ResponseWriter, r *http.Request){
}

func SuperSecure(c *web.C, h http.Handler) http.Handler {
	fn := func (w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Basic ") {
			pleaseAuth(w)
			return
		}

		password, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil || string(password) != Password {
			pleaseAuth(w)
			return
		}

	}
	return http.HandlerFunc(fn)
}

func pleaseAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Gritter"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Go away!\n"))
}