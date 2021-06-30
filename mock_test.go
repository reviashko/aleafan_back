package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/reviashko/aleafan_back/models"

	"github.com/lib/pq"
)

type mockDB struct{}

func (mdb *mockDB) GetQuestions(employeeid string) ([]*models.Question, pq.ErrorCode, error) {
	questionList := make([]*models.Question, 0)
	questionList = append(questionList, &models.Question{QuestionID: 35, QuestionText: "Вы всегда контролируете свое поведение?"})
	questionList = append(questionList, &models.Question{QuestionID: 27, QuestionText: "Можно ли Вас назвать гурманом?"})

	return questionList, "", nil
}

func (mdb *mockDB) GetQuestionsJSON(employeeid string) (string, pq.ErrorCode, error) {

	questionList, erroCode, err := mdb.GetQuestions(employeeid)

	e, err := json.Marshal(questionList)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return string(e), erroCode, nil
}

func (mdb *mockDB) GetSurveyResult(employeeid string) (*models.SurveyResult, pq.ErrorCode, error) {

	//models.SurveyResult
	surveyResult := new(models.SurveyResult)
	surveyResult.AfanasievType = "Lenin"
	surveyResult.Descr = "123"
	surveyResult.FeedBackExists = 50
	surveyResult.TypeName = "VLFE"

	return surveyResult, "", nil
}

func (mdb *mockDB) SaveAnswers(employeeid string, answers string) (pq.ErrorCode, error) {
	//TODO make real save
	return "", nil
}

// SaveSurveyFeedBack method
func (mdb *mockDB) SaveSurveyFeedBack(employeeid string, surveyfeedback int) error {

	return nil
}

//TestGetQuestionsJSON func
func TestGetQuestionsJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.getQuestionJSONHandler).ServeHTTP(rec, req)

	expected := `[{"qst_id":35,"qst_text":"Вы всегда контролируете свое поведение?"},{"qst_id":27,"qst_text":"Можно ли Вас назвать гурманом?"}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
