package models

import (
	"github.com/lib/pq"
)

// SurveyResult structure
type SurveyResult struct {
	AfanasievType  string `json:"afanasiev_type" db:"afanasiev_type"`
	Descr          string `json:"ps_descr" db:"ps_descr"`
	TypeName       string `json:"ps_name" db:"ps_name"`
	FeedBackExists bool   `json:"afanasiev_feedback" db:"afanasiev_feedback"`
}

// GetSurveyResult method
func (db *DB) GetSurveyResult(employeeid int) (*SurveyResult, pq.ErrorCode, error) {

	surveyResult := new(SurveyResult)
	var errorCode pq.ErrorCode

	rows, err := db.Queryx("select afanasiev_type, ps_descr, ps_name, afanasiev_feedback from testing.employee_getresult($1)", employeeid)
	defer rows.Close()
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			errorCode = err.Code
		}
		return surveyResult, errorCode, err
	}

	for rows.Next() {
		err := rows.StructScan(&surveyResult)
		if err != nil {
			return surveyResult, errorCode, err
		}
	}

	return surveyResult, errorCode, nil
}

// SaveSurveyFeedBack method
func (db *DB) SaveSurveyFeedBack(employeeid int, surveyfeedback int) error {

	rows, err := db.Queryx("select * from testing.feedback_save($1, $2)", employeeid, surveyfeedback)
	defer rows.Close()
	if err != nil {
		return err
	}

	return nil
}
