package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var BASE_URLP = "http://localhost:9000/api/v1/applicantusers"

type ApplicantCreateImp struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Gmail    string `json:"gmail"`
	Telepon  int    `json:"telepon"`
	Age      int    `json:"age"`
	Job      string `json:"job"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
	Name     string `json:"name"`
}

func (a ApplicantCreateImp) Store(w http.ResponseWriter, r *http.Request) {

}

func (a ApplicantCreateImp) Create(w http.ResponseWriter, r *http.Request) {
	var post ApplicantCreateImp
	var data map[string]interface{}

	id := r.URL.Query().Get("id")

	fmt.Println(id)

	if id != "" {
		res, err := http.Get(BASE_URLP + "/applicants/")
		if err != nil {
			log.Print(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(res.Body)

		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&post); err != nil {
			log.Print(err)
		}

		data = map[string]interface{}{
			"post": post,
		}
	}

	temp, _ := template.ParseFiles("")
	err := temp.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (a ApplicantCreateImp) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicantCreateImp) Index(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
