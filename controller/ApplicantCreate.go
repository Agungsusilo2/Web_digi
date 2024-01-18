package controller

import "net/http"

type ApplicantCreate interface {
	Store(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}
