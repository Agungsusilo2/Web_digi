package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	dto "DigiHero/DTO"
)

const (
	apiKey   = "Pukulele"
	base_url = "http://localhost:9888/api/v1"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("view/home.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)

}

func RegHeroes(w http.ResponseWriter, r *http.Request) {

	var post dto.Heroes
	var data map[string]interface{}

	data = map[string]interface{}{
		"post": post,
	}

	temp, _ := template.ParseFiles("view/Heroes.html")
	temp.Execute(w, data)

}

func ActRegHeroes(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	newApp := dto.Heroes{
		NamaLengkap:  r.Form.Get("AppFullName"),
		Email:        r.Form.Get("AppEmail"),
		NomorTelepon: r.Form.Get("AppNumber"),
		Umur:         r.Form.Get("AppAge"),
		Pekerjaan:    r.Form.Get("AppJob"),
	}

	jsonValue, _ := json.Marshal(newApp)
	buff := bytes.NewBuffer(jsonValue)

	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodPost, base_url+"/applicantusers", buff)

	if err != nil {
		log.Print(err)
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {

		var NewIdApp struct {
			Id string `json:"Id"`
		}

		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&NewIdApp); err != nil {
			log.Print(err)
		}

		newAppID := NewIdApp.Id
		fmt.Println(newAppID)

		if res.StatusCode == 201 || res.StatusCode == 200 {
			http.Redirect(w, r, "/NextStep/Heroes?Id="+newAppID, http.StatusSeeOther)
		}

	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	// fmt.Println(postResponse)

	// if res.StatusCode == 201 || res.StatusCode == 200 {
	// 	http.Redirect(w, r, "/Succes", http.StatusSeeOther)
	// }

}

func NextHeroes(w http.ResponseWriter, r *http.Request) {
	var post dto.NextHeroes
	var data map[string]interface{}

	data = map[string]interface{}{
		"post": post,
	}

	temp, _ := template.ParseFiles("view/Heroes1.html")
	temp.Execute(w, data)
}

func NextRegHeroes(w http.ResponseWriter, r *http.Request) {

	// var post Post
	// var data map[string]interface{}

	ID := r.URL.Query().Get("Id")

	NextReg := dto.NextHeroes{
		Id:        ID,
		Countries: r.Form.Get("Countries"),
		Skills:    r.Form.Get("Skills"),
		Username:  r.Form.Get("Username"),
		Password:  r.Form.Get("Password"),
	}

	jsonValue, _ := json.Marshal(NextReg)
	buff := bytes.NewBuffer(jsonValue)

	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodPatch, base_url+"/applicantusers/NextReg"+ID, buff)

	if err != nil {
		log.Print(err)
	}

	if err != nil {
		log.Print(err)
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	var postResponse dto.NextHeroes

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&postResponse); err != nil {
		log.Print(err)
	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(postResponse)

	if res.StatusCode == 201 || res.StatusCode == 200 {
		http.Redirect(w, r, "/Succes", http.StatusSeeOther)
	}
}

func RegSuccess(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("view/RegistrasiBerhasil.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)

}

func Started(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("view/login.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)

}
