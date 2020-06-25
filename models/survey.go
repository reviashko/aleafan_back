package models

import (
	"encoding/json"
	//"fmt"

	"github.com/lib/pq"
)

// Question structure
type Question struct {
	QuestionID   int    `json:"qst_id" db:"qst_id"`
	QuestionText string `json:"qst_text" db:"qst_text"`
}

// GetQuestionsJSON method
func (db *DB) GetQuestionsJSON(employeeid int) (string, pq.ErrorCode, error) {

	questionList, errorCode, err := db.GetQuestions(employeeid)
	if err != nil {
		return "", errorCode, err
	}

	e, err := json.Marshal(questionList)
	if err != nil {
		return "", errorCode, err
	}

	return string(e), errorCode, nil
}

// GetQuestions method
func (db *DB) GetQuestions(employeeid int) ([]*Question, pq.ErrorCode, error) {

	questionList := make([]*Question, 0)
	var errorCode pq.ErrorCode

	rows, err := db.Queryx("select qst_id, qst_text from testing.questions_getall($1)", employeeid)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			errorCode = err.Code
		}
		return nil, errorCode, err
	}
	defer rows.Close()

	for rows.Next() {
		question := new(Question)
		err := rows.StructScan(&question)
		if err != nil {
			return nil, errorCode, err
		}
		questionList = append(questionList, question)
	}

	return questionList, errorCode, nil
}

// SaveAnswers method
func (db *DB) SaveAnswers(employeeid int, answers string) (pq.ErrorCode, error) {

	var errorCode pq.ErrorCode

	rows, err := db.Queryx("select * from testing.testing_save($1, $2)", employeeid, answers)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			errorCode = err.Code
		}
		return errorCode, err
	}
	defer rows.Close()

	return errorCode, nil
}
