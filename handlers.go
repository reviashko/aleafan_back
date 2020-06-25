package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (env *Env) surveyResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	vars := mux.Vars(r)
	employeeid, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data, _, err := env.db.GetSurveyResult(employeeid)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", data)
}

func (env *Env) getQuestionJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	vars := mux.Vars(r)
	employeeid, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data, errorCode, err := env.db.GetQuestionsJSON(employeeid)
	if err != nil {
		if errorCode == "22024" { //answered earlyer
			//http.Error(w, http.StatusText(301), 301)

			data, _, _ = env.db.GetSurveyResult(employeeid)

		} else {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", data)
}

func (env *Env) saveSurveyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	bodyString := string(body)

	type Answer struct {
		QstID  int `json:"qst_id"`
		Answer int `json:"answer"`
	}

	type Poll struct {
		Answers []Answer `json:"answers"`
		UserID  int      `json:"userid"`
	}

	var p Poll
	err = json.Unmarshal([]byte(bodyString), &p)
	if err != nil {
		fmt.Println(err)
	}

	answers, err := json.Marshal(p.Answers)

	errCode, err := env.db.SaveAnswers(p.UserID, string(answers))
	if err != nil && errCode != "22024" {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data, _, err := env.db.GetSurveyResult(p.UserID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", data)
}

func (env *Env) saveSurveyFeedBack(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	bodyString := string(body)

	type FeedBackAnswer struct {
		Answer int `json:"answer"`
		UserID int `json:"userid"`
	}

	var fba FeedBackAnswer
	err = json.Unmarshal([]byte(bodyString), &fba)
	if err != nil {
		fmt.Println(err)
	}

	err = env.db.SaveSurveyFeedBack(fba.UserID, fba.Answer)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", "{'result':'ok'}")
}
