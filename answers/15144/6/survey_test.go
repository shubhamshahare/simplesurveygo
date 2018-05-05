package tests_test

import (
	"simplesurveygo/dao"
	"testing"
        "time"
        "fmt"
)

func TestCreateQuestions(t *testing.T) {
	session := dao.MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("survey")

	q1 := dao.Question{QuestionString: "What age bracket do you fall in?",
		Options: []string{"<18", "18-30", "30-45", "45+"}}
	q2 := dao.Question{QuestionString: "Do you have a job?",
		Options: []string{"Yes", "No"}}
	q3 := dao.Question{QuestionString: "Do you have a stable source of income?",
		Options: []string{"Yes", "No", "Sometimes"}}
	q4 := dao.Question{QuestionString: "Would you consider getting a credit card?",
		Options: []string{"Definitely", "Maybe", "Probably not", "Definitely Not"}}

	questions := []dao.Question{q1, q2, q3, q4}
        surveyduration := int(5)
        surveystarttime := time.Now()
        fmt.Println(surveyduration)
	survey := dao.Survey{SurveyName: "creditcard",
		Heading:     "Consider a credit card!",
		Description: "Checking the chances of people option for a credit card",
		Questions:   questions,
		Status:      true ,
                SurveyDuration: surveyduration,
                SurveyStartTime: surveystarttime ,
             }


	clctn.Insert(survey)

	q1 = dao.Question{QuestionString: "Which state do you belong to?",
		Options: []string{"Maharashtra", "Gujarat", "Rajasthan", "Other"}}
	q2 = dao.Question{QuestionString: "When was the last time you had coffe?",
		Options: []string{"Been years", "A few months", "Just a few weeks", "Yesterday", "Having one right now"}}
	q3 = dao.Question{QuestionString: "What king do you have?",
		Options: []string{"Regular", "Decaf", "Expresso"}}
	q4 = dao.Question{QuestionString: "Do you have coffe with sugar?",
		Options: []string{"Yes", "No", "Sometimes"}}
	q5 := dao.Question{QuestionString: "Which brand do you prefer?",
		Options: []string{"Nescafe", "Girnar", "Bru"}}

	questions = []dao.Question{q1, q2, q3, q4, q5}

	survey = dao.Survey{SurveyName: "coffelovers",
		Heading:     "We love coffe!",
		Description: "Checking the love people have for coffe",
		Questions:   questions,
		Status:      false,
                SurveyDuration: surveyduration,
                SurveyStartTime: surveystarttime ,
              }

	clctn.Insert(survey)

	q1 = dao.Question{QuestionString: "Are you a Guy or a Girl",
		Options: []string{"Guy", "Girl"}}
	q2 = dao.Question{QuestionString: "Do you like music?",
		Options: []string{"Yes", "Occasionally", "No"}}
	q3 = dao.Question{QuestionString: "What kind of music do you prefer?",
		Options: []string{"Classical", "Pop", "Metal", "Rock"}}

	questions = []dao.Question{q1, q2, q3}

	survey = dao.Survey{SurveyName: "musiclovers",
		Heading:     "We love music!",
		Description: "Checking the love people have for music",
		Questions:   questions,
		Status:      true,
                SurveyDuration: surveyduration,
                SurveyStartTime: surveystarttime ,
              }
                

	clctn.Insert(survey)
}

