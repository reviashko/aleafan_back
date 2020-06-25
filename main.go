package main

import (
	"aleafan/models"
	"log"
	"net/http"

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

	router.HandleFunc("/getquestionjson/{id}", env.getquestionjsonhandler).Methods("GET")
	router.HandleFunc("/surveyresult/{id}", env.surveyresulthandler).Methods("GET")
	router.HandleFunc("/savesurveyfeedback/", env.savesurveyfeedback).Methods("POST")
	router.HandleFunc("/savesurvey/", env.savesurveyhandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
