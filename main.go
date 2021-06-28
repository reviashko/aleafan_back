package main

import (
	"log"
	"net/http"

	"github.com/reviashko/aleafan_back/models"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

//Env struct
type Env struct {
	db models.Datastore
}

func main() {

	connectionData := models.ConnectionData{}
	if gonfig.GetConf("config/db.json", &connectionData) != nil {
		log.Panic("load confg error")
	}

	db, err := models.InitDB(connectionData.ToString())
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db}

	router := mux.NewRouter()

	router.HandleFunc("/getquestionjson/{id}", env.getQuestionJSONHandler).Methods("GET")
	router.HandleFunc("/surveyresult/{id}", env.surveyResultHandler).Methods("GET")
	router.HandleFunc("/savesurveyfeedback/", env.saveSurveyFeedBack).Methods("POST")
	router.HandleFunc("/savesurvey/", env.saveSurveyHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
